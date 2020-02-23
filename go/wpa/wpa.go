package wpa

import (
	"log"
	"net"
	"strings"
	"time"

	wpasupplicant "github.com/dpifke/golang-wpasupplicant"
	"github.com/pkg/errors"
)

type Wpa struct {
	Conn wpasupplicant.Conn

	interfaceName string
}

func (w *Wpa) startConnection() error {
	wpa, err := wpasupplicant.Unixgram(w.interfaceName)
	if err != nil {
		return errors.Wrapf(err, "could not init wpasupplicant for %s", w.interfaceName)
	}

	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				return
			case msg := <-wpa.EventQueue():
				log.Printf("Wpa event: %v", msg)
			}
		}
	}()

	go func() {
		defer func() {
			stop <- true
		}()
		for range time.Tick(time.Second) {
			if err := wpa.Ping(); err != nil {
				log.Printf("Ping failed: %v, restarting connection.", err)
				if err := w.startConnection(); err != nil {
					log.Printf("Could not restart connection: %v", err)
					continue
				}
				return
			}
		}
	}()

	w.Conn = wpa

	return nil
}

func InitWpa() (*Wpa, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.Wrap(err, "could not list network interfaces")
	}

	name := ""
	for _, intf := range interfaces {
		if strings.HasPrefix(intf.Name, "w") {
			name = intf.Name
			break
		}
	}

	if name == "" {
		return nil, errors.New("no wifi network interface found")
	}

	wpa := &Wpa{
		interfaceName: name,
	}

	return wpa, wpa.startConnection()
}
