package bluetooth

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type connHandler struct {
	conn   net.Conn
	server *Server

	wifiSenderStop chan struct{}
}

func (h *connHandler) Update(game *chess.Game, move *chess_state.Move) error {
	msg := &bluetoothpb.Response{
		ChessBoard: &bluetoothpb.Response_ChessBoard{
			Fen: game.FEN(),
		},
	}

	return writeProto(h.conn, msg)
}

func (h *connHandler) Handle() error {
	for {
		request := &bluetoothpb.Request{}
		if err := readProto(h.conn, request); err != nil {
			if strings.Contains(err.Error(), "connection reset by peer") {
				return nil
			}
			return errors.Wrap(err, "could not read proto")
		}

		if err := h.handleRequest(request); err != nil {
			return err
		}
	}
}

func (h *connHandler) handleRequest(request *bluetoothpb.Request) error {
	h.server.mutex.Lock()
	defer h.server.mutex.Unlock()

	switch request.GetType() {
	case bluetoothpb.Request_CONFIGURE_WIFI:
		var networkID int
		network, err := h.server.getNetwork(request.GetWifiSsid())
		if err != nil && strings.Contains(err.Error(), "not found") {
			networkID, err = h.server.Wpa.AddNetwork()
			if err != nil {
				return errors.Wrap(err, "could not add new network")
			}
		} else if err != nil {
			return err
		} else {
			networkID, err = strconv.Atoi(network.NetworkID())
			if err != nil {
				return errors.Wrapf(err, "could not parse network id '%v'", network.NetworkID())
			}
		}
		if err := h.server.Wpa.SetNetwork(networkID, "ssid", request.GetWifiSsid()); err != nil {
			return errors.Wrap(err, "could not set network ssid")
		}
		if err := h.server.Wpa.SetNetwork(networkID, "psk", request.GetWifiPsk()); err != nil {
			return errors.Wrap(err, "could not set network psk")
		}
		if err := h.server.Wpa.SaveConfig(); err != nil {
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
			return errors.Wrapf(err, "could not parse network id '%v'", network.NetworkID())
		}
		if err := h.server.Wpa.SelectNetwork(networkID); err != nil {
			return errors.Wrapf(err, "could not connect to %v", network.SSID())
		}
	case bluetoothpb.Request_FORGET_WIFI:
		network, err := h.server.getNetwork(request.GetWifiSsid())
		if err != nil {
			return err
		}
		networkID, err := strconv.Atoi(network.NetworkID())
		if err != nil {
			return errors.Wrapf(err, "could not parse network id '%v'", network.NetworkID())
		}
		if err := h.server.Wpa.RemoveNetwork(networkID); err != nil {
			return errors.Wrap(err, "could not remove network")
		}
		if err := h.server.Wpa.SaveConfig(); err != nil {
			return errors.Wrap(err, "could not save wpa config")
		}
	case bluetoothpb.Request_START_WIFI_SCAN:
		if err := h.server.Wpa.Scan(); err != nil {
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

	response := &bluetoothpb.Response{}

	configuredNetworks, err := h.server.Wpa.ListNetworks()
	if err != nil {
		return errors.Wrap(err, "could not list networks")
	}
	for _, network := range configuredNetworks {
		response.ConfiguredNetworks = append(response.ConfiguredNetworks, &bluetoothpb.Response_WifiNetwork{
			Ssid: network.SSID(),
		})
	}
	discoveredNetworks, errs := h.server.Wpa.ScanResults()
	if errs != nil {
		return errors.Errorf("could not list scan results: %v", errs)
	}
	for _, network := range discoveredNetworks {
		response.DiscoveredNetworks = append(response.DiscoveredNetworks, &bluetoothpb.Response_WifiNetwork{
			Ssid: network.SSID(),
		})
	}

	if err := writeProto(h.conn, response); err != nil {
		return errors.Wrap(err, "could not write response")
	}

	return nil
}
