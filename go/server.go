package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type Server struct {
	Channel   int
	Observers *Observers

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
	s.Observers.Add(&bluetoothObserver{conn})

	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			if strings.Contains(err.Error(), "connection reset by peer") {
				return
			}
			log.Println(err)
			return
		}

		log.Println(string(data[:n]))

	}
}

func writeProto(w io.Writer, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "could not marshal proto message")
	}

	log.Println(len(data))
	if err := binary.Write(w, binary.BigEndian, uint64(len(data))); err != nil {
		return errors.Wrap(err, "could not write data length")
	}

	_, err = w.Write(data)
	if err != nil {
		return errors.Wrap(err, "could not write data")
	}

	return nil
}

type bluetoothObserver struct {
	conn net.Conn
}

func (o *bluetoothObserver) Update(game *chess.Game) error {
	msg := &bluetoothpb.Response{
		ChessBoard: &bluetoothpb.Response_ChessBoard{
			Fen: game.FEN(),
		},
	}

	return writeProto(o.conn, msg)
}
