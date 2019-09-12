package bluetooth

import (
	"fmt"
	"log"
	"net"
	"sync"

	wpasupplicant "github.com/dpifke/golang-wpasupplicant"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/pkg/errors"
)

type Server struct {
	Channel   int
	Observers *chess_state.Observers

	mutex sync.Mutex
	Wpa   wpasupplicant.Conn

	ln net.Listener
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
	s.Observers.Add(handler)

	if err := handler.Handle(); err != nil {
		log.Println(err)
	}
}
func (s *Server) getNetwork(ssid string) (wpasupplicant.ConfiguredNetwork, error) {
	networks, err := s.Wpa.ListNetworks()
	if err != nil {
		return nil, errors.Wrap(err, "could not list wifi networks")
	}
	for _, network := range networks {
		if network.SSID() == ssid {
			return network, nil
		}
	}

	return nil, errors.Errorf("network '%v' not found", ssid)
}
