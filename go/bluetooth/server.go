package bluetooth

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	wpasupplicant "github.com/dpifke/golang-wpasupplicant"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/matematik7/dichess/go/wpa"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type Server struct {
	Channel    int
	Controller *chess_state.Controller

	mutex sync.Mutex
	Wpa   *wpa.Wpa

	moveChan  chan string
	startChan chan bool

	ln net.Listener
}

func NewServer(channel int, controller *chess_state.Controller, wpa *wpa.Wpa) *Server {
	return &Server{
		Channel:    channel,
		Controller: controller,
		Wpa:        wpa,

		moveChan:  make(chan string),
		startChan: make(chan bool),
	}
}

func (s *Server) Serve() error {
	var err error
	s.ln, err = rfcomm.Listen(fmt.Sprintf(":%d", s.Channel))
	if err != nil {
		return err
	}
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go s.Handle(conn)
	}
}

func (s *Server) OnShutdown(string) error {
	return s.ln.Close()
}

func (s *Server) Handle(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	handler := &connHandler{conn: conn, server: s}
	observerId := s.Controller.Observers.Add(handler)
	defer s.Controller.Observers.Remove(observerId)
	stateSenderId := s.Controller.StateSenders.Add(handler)
	defer s.Controller.StateSenders.Remove(stateSenderId)

	if err := handler.Handle(); err != nil {
		log.Printf("Handler exited: %v", err)
		return
	}
}
func (s *Server) getNetwork(ssid string) (wpasupplicant.ConfiguredNetwork, error) {
	networks, err := s.Wpa.Conn.ListNetworks()
	if err != nil {
		return nil, errors.Wrap(err, "could not list wifi networks")
	}
	for _, network := range networks {
		if network.SSID() == ssid {
			return network, nil
		}
	}

	return nil, errors.Errorf("network '%s' not found", ssid)
}

func (s *Server) MakeMove(ctx context.Context, stateSender chess_state.StateSender, game *chess.Game) (*chess_state.Move, error) {
	move := &chess_state.Move{
		ShouldMove: true,
		ShouldSay:  true,
	}
	select {
	case <-ctx.Done():
		return move, nil
	case moveStr := <-s.moveChan:
		if moveStr == "UNDO" {
			move.Undo = true
			return move, nil
		}
		mv, err := chess.AlgebraicNotation{}.Decode(game.Position(), moveStr)
		if err != nil {
			return nil, err
		}
		move.Move = mv
		return move, nil
	}
}

func (s *Server) StartGame(chess_state.StateSender) error {
	select {
	case <-s.startChan:
		return nil
	}
}
