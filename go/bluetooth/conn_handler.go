package bluetooth

import (
	"context"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type connHandler struct {
	mutex  sync.Mutex
	conn   net.Conn
	server *Server

	wifiSenderStop chan struct{}
}

func (h *connHandler) Update(_ context.Context, _ chess_state.StateSender, game *chess_state.Game, move *chess_state.Move) error {
	return h.sendUpdate(game, move, nil)
}

func (h *connHandler) sendUpdate(game *chess_state.Game, move *chess_state.Move, settings *bluetoothpb.Settings) error {
	msg := &bluetoothpb.Response{
		Type:           bluetoothpb.Response_GAME_UPDATE,
		GameInProgress: game != nil,
		Settings:       settings,
	}
	if game != nil {
		rotate := false
		if _, ok := game.Players[1].(*chess_state.HumanPlayer); ok {
			rotate = true
			if game.Game.Position().Turn() == chess.White {
				if _, ok := game.Players[0].(*chess_state.HumanPlayer); ok {
					rotate = false
				}
			}
		}

		notation := chess.AlgebraicNotation{}
		positions := game.Game.Positions()
		for i, move := range game.Game.Moves() {
			msg.Moves = append(msg.Moves, notation.Encode(positions[i], move))
		}

		canMakeMove := false
		if game.Game.Position().Turn() == chess.White {
			if _, ok := game.Players[0].(*chess_state.HumanPlayer); ok {
				canMakeMove = true
			}
		} else {
			if _, ok := game.Players[1].(*chess_state.HumanPlayer); ok {
				canMakeMove = true
			}
		}

		if game.Game.Outcome() != chess.NoOutcome {
			canMakeMove = false
			if len(msg.Moves)%2 == 1 {
				msg.Moves = append(msg.Moves, "")
			}
			msg.Moves = append(msg.Moves, game.Game.Outcome().String())
		}

		msg.ChessBoard = &bluetoothpb.Response_ChessBoard{
			Fen:         game.Game.FEN(),
			Rotate:      rotate,
			CanMakeMove: canMakeMove,
		}
	}

	return h.send(msg)
}

func (h *connHandler) StateSend(state string) {
	msg := &bluetoothpb.Response{
		Type:  bluetoothpb.Response_STATE_UPDATE,
		State: state,
	}
	if err := h.send(msg); err != nil {
		log.Printf("Could not send state: %v", err)
	}
}

func (h *connHandler) send(msg proto.Message) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	return writeProto(h.conn, msg)
}

func (h *connHandler) Handle() error {
	time.Sleep(100 * time.Millisecond)
	if err := h.sendUpdate(h.server.Controller.GetGame(), nil, h.server.Controller.GetSettings()); err != nil {
		return err
	}
	h.StateSend(h.server.Controller.StateSenders.GetLastState())
	for {
		request := &bluetoothpb.Request{}
		if err := readProto(h.conn, request); err != nil {
			if strings.Contains(err.Error(), "connection reset by peer") {
				return nil
			}
			return errors.Wrap(err, "could not read proto")
		}

		if err := h.handleRequest(request); err != nil {
			log.Printf("Could not handle request: %v", err)
		}
	}
}

func (h *connHandler) handleRequest(request *bluetoothpb.Request) error {
	h.server.mutex.Lock()
	defer h.server.mutex.Unlock()

	log.Printf("Got request: %v", request.GetType())

	switch request.GetType() {
	case bluetoothpb.Request_CONFIGURE_WIFI:
		var networkID int
		network, err := h.server.getNetwork(request.GetWifiSsid())
		if err != nil && strings.Contains(err.Error(), "not found") {
			networkID, err = h.server.Wpa.Conn.AddNetwork()
			if err != nil {
				return errors.Wrap(err, "could not add new network")
			}
		} else if err != nil {
			return err
		} else {
			networkID, err = strconv.Atoi(network.NetworkID())
			if err != nil {
				return errors.Wrapf(err, "could not parse network id '%s'", network.NetworkID())
			}
		}
		if err := h.server.Wpa.Conn.SetNetwork(networkID, "ssid", request.GetWifiSsid()); err != nil {
			return errors.Wrap(err, "could not set network ssid")
		}
		if err := h.server.Wpa.Conn.SetNetwork(networkID, "psk", request.GetWifiPsk()); err != nil {
			return errors.Wrap(err, "could not set network psk")
		}
		if err := h.server.Wpa.Conn.SaveConfig(); err != nil {
			return errors.Wrap(err, "could not save wpa config")
		}
		// continue to connect after configure
		fallthrough
	case bluetoothpb.Request_CONNECT_WIFI:
		network, err := h.server.getNetwork(request.GetWifiSsid())
		if err != nil {
			return err
		}
		networkID, err := strconv.Atoi(network.NetworkID())
		if err != nil {
			return errors.Wrapf(err, "could not parse network id '%s'", network.NetworkID())
		}
		if err := h.server.Wpa.Conn.SelectNetwork(networkID); err != nil {
			return errors.Wrapf(err, "could not connect to %s", network.SSID())
		}
	case bluetoothpb.Request_FORGET_WIFI:
		network, err := h.server.getNetwork(request.GetWifiSsid())
		if err != nil {
			return err
		}
		networkID, err := strconv.Atoi(network.NetworkID())
		if err != nil {
			return errors.Wrapf(err, "could not parse network id '%s'", network.NetworkID())
		}
		if err := h.server.Wpa.Conn.RemoveNetwork(networkID); err != nil {
			return errors.Wrap(err, "could not remove network")
		}
		if err := h.server.Wpa.Conn.SaveConfig(); err != nil {
			return errors.Wrap(err, "could not save wpa config")
		}
	case bluetoothpb.Request_START_WIFI_SCAN:
		if err := h.server.Wpa.Conn.Scan(); err != nil {
			return errors.Wrap(err, "could not start scan")
		}
		if h.wifiSenderStop == nil {
			h.wifiSenderStop = make(chan struct{})
			go h.wifiSender()
		}

	case bluetoothpb.Request_STOP_WIFI_SCAN:
		if h.wifiSenderStop != nil {
			close(h.wifiSenderStop)
			h.wifiSenderStop = nil
		}

	case bluetoothpb.Request_UPDATE_SETTINGS:
		h.server.Controller.SetSettings(request.GetSettings())
		if err := h.sendUpdate(h.server.Controller.GetGame(), nil, h.server.Controller.GetSettings()); err != nil {
			return err
		}
	case bluetoothpb.Request_MOVE:
		select {
		case h.server.moveChan <- request.Move:
		default:
		}
	case bluetoothpb.Request_UNDO_MOVE:
		select {
		case h.server.moveChan <- "UNDO":
		default:
		}
	case bluetoothpb.Request_NEW_GAME:
		h.server.Controller.StopGame()
		if err := h.sendUpdate(nil, nil, h.server.Controller.GetSettings()); err != nil {
			return err
		}
	case bluetoothpb.Request_GET_SETTINGS:
		if err := h.sendUpdate(h.server.Controller.GetGame(), nil, h.server.Controller.GetSettings()); err != nil {
			return err
		}
	case bluetoothpb.Request_START_GAME:
		select {
		case h.server.startChan <- true:
		default:
		}
	}

	return nil
}

func (h *connHandler) wifiSender() {
	stop := h.wifiSenderStop
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		if err := h.sendWifiInfo(); err != nil {
			log.Printf("Could  not send wifi info: %v", err)
		}

		select {
		case <-ticker.C:
		case <-stop:
			return
		}
	}
}

func (h *connHandler) sendWifiInfo() error {
	h.server.mutex.Lock()
	defer h.server.mutex.Unlock()

	networks := make(map[string]*bluetoothpb.Response_WifiNetwork)

	configuredNetworks, err := h.server.Wpa.Conn.ListNetworks()
	if err != nil {
		return errors.Wrap(err, "could not list networks")
	}
	for _, network := range configuredNetworks {
		networks[network.SSID()] = &bluetoothpb.Response_WifiNetwork{
			Ssid:       network.SSID(),
			Saved:      true,
			Connecting: true,
		}
		for _, flag := range network.Flags() {
			if flag == "CURRENT" {
				networks[network.SSID()].Connected = true
				networks[network.SSID()].Connecting = false
			}
			if flag == "DISABLED" {
				networks[network.SSID()].Connecting = false
			}
			if flag == "TEMP-DISABLED" {
				networks[network.SSID()].Connecting = false
				networks[network.SSID()].Failed = true
			}
		}
	}
	discoveredNetworks, errs := h.server.Wpa.Conn.ScanResults()
	if errs != nil {
		return errors.Errorf("could not list scan results: %v", errs)
	}
	for _, network := range discoveredNetworks {
		if network.SSID() == "" {
			continue
		}
		if _, ok := networks[network.SSID()]; ok {
			networks[network.SSID()].Available = true
		} else {
			networks[network.SSID()] = &bluetoothpb.Response_WifiNetwork{
				Ssid:      network.SSID(),
				Available: true,
			}
		}
	}

	networksList := make([]*bluetoothpb.Response_WifiNetwork, 0, len(networks))
	for _, network := range networks {
		networksList = append(networksList, network)
	}
	sort.Slice(networksList, func(i, j int) bool {
		if networksList[i].Connected != networksList[j].Connected {
			return networksList[i].Connected
		}
		if networksList[i].Available != networksList[j].Available {
			return networksList[i].Available
		}
		if networksList[i].Saved != networksList[j].Saved {
			return networksList[i].Saved
		}
		return strings.Compare(networksList[i].Ssid, networksList[j].Ssid) < 0
	})

	response := &bluetoothpb.Response{
		Type:     bluetoothpb.Response_WIFI_UPDATE,
		Networks: networksList,
	}

	if err := writeProto(h.conn, response); err != nil {
		return errors.Wrap(err, "could not write response")
	}

	return nil
}
