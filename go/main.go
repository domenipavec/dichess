package main

import (
	"context"
	"log"

	"github.com/Zemanta/gracefulshutdown"
	"github.com/Zemanta/gracefulshutdown/shutdownmanagers/posixsignal"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/muka/go-bluetooth/api"
	"github.com/pkg/errors"
)

const channel = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Make sure periph is initialized.
	// if _, err := host.Init(); err != nil {
	//     log.Fatal(err)
	// }
	//
	// matrix, err := hardware.NewMatrix()
	// if err != nil {
	//     log.Fatal(err)
	// }
	// for {
	//     data, err := matrix.Read()
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     log.Println("start")
	//     for _, row := range data {
	//         log.Println(row)
	//     }
	//     time.Sleep(time.Second)
	// }
	// return
	// if err := voiceRecognition(ctx); err != nil {
	//     log.Fatal(err)
	// }
	// return

	observers := &Observers{}
	observers.Add(&LoggingObserver{})
	server := &Server{
		Channel:   channel,
		Observers: observers,
	}

	gs := gracefulshutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	// dichessProfile := rfcomm.NewSerialProfile("dichess", "4f067110-8c71-488a-abf1-5606375d0dd8", channel)
	dichessProfile := rfcomm.NewSerialProfile("dichess", "00001101-0000-1000-8000-00805f9b34fb", channel)
	// androidAutoProfile := rfcomm.NewSerialProfile("androidauto", "4de17a00-52cb-11e6-bdf4-0800200c9a66", channel)

	err := dichessProfile.Register()
	if err != nil {
		log.Fatalf("Could not register android audo profile: %v", err)
	}
	gs.AddShutdownCallback(dichessProfile)
	gs.AddShutdownCallback(server)

	if err := gs.Start(); err != nil {
		log.Fatalf("Could not start graceful shutdown: %v", err)
	}

	if err := startBluetoothController(ctx); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := newGame(observers); err != nil {
			log.Println(err)
		}
	}()

	log.Fatal(server.Serve())
}

func newGame(observers *Observers) error {
	player1, err := NewUciPlayer()
	if err != nil {
		return err
	}
	player2, err := NewUciPlayer()
	if err != nil {
		return err
	}
	cs := NewChessState(player1, player2, observers)
	go func() {
		if err := cs.Play(); err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func startBluetoothController(ctx context.Context) error {
	// cmd := exec.CommandContext(ctx, "btattach", "-N", "-S", "115200", "-P", "bcm", "-B", "/dev/ttyAMA0")
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// err := cmd.Start()
	// if err != nil {
	//     return errors.Wrap(err, "could not start btattach")
	// }
	//
	// go func() {
	//     err := cmd.Wait()
	//     log.Printf("btattach exited with: %v", err)
	// }()
	//
	// for i := 0; i < 60; i++ {
	//     time.Sleep(time.Second)
	//     exists, err := api.AdapterExists("hci0")
	//     if err != nil {
	//         return errors.Wrap(err, "could not check for adapter hci0")
	//     }
	//     if exists {
	//         break
	//     }
	// }

	adapter, err := api.GetAdapter("hci0")
	if err != nil {
		return errors.Wrapf(err, "couldn't find adapter %v", "hci0")
	}

	if err := adapter.SetProperty("Powered", true); err != nil {
		return errors.Wrap(err, "couldn't set powered to true")
	}
	if err := adapter.SetProperty("Alias", "dichess"); err != nil {
		return errors.Wrap(err, "couldn't set powered to true")
	}
	if err := adapter.SetProperty("DiscoverableTimeout", uint32(0)); err != nil {
		return errors.Wrap(err, "couldn't set discoverable timeout to 0")
	}
	if err := adapter.SetProperty("Discoverable", true); err != nil {
		return errors.Wrap(err, "couldn't set discoverable to true")
	}

	return nil
}
