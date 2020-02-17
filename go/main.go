package main

import (
	"context"
	"flag"
	"log"

	"github.com/Zemanta/gracefulshutdown"
	"github.com/Zemanta/gracefulshutdown/shutdownmanagers/posixsignal"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/matematik7/dichess/go/bluetooth"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/matematik7/dichess/go/hardware"
	"github.com/matematik7/dichess/go/voice"
	"github.com/matematik7/dichess/go/wpa"
)

var noHardware = flag.Bool("no_hardware", false, "disable hardware init and use fake")

const btChannel = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	hw := hardware.New()
	if *noHardware {
		if err := hw.InitializeFake(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := hw.InitializeReal(); err != nil {
			log.Fatal(err)
		}
	}

	observers := &chess_state.Observers{}
	observers.Add(&chess_state.LoggingObserver{})
	observers.Add(hw)

	voice, err := voice.New(ctx)
	if err != nil {
		log.Printf("Couldn't init voice: %v", err)
		voice = nil
	} else {
		observers.Add(voice)
	}

	wpa, err := wpa.InitWpa()
	if err != nil {
		log.Fatal(err)
	}

	controller := &chess_state.Controller{Observers: observers}
	server := &bluetooth.Server{
		Controller: controller,
		Channel:    btChannel,
		Wpa:        wpa,
	}

	gs := gracefulshutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	// dichessProfile := rfcomm.NewSerialProfile("dichess", "4f067110-8c71-488a-abf1-5606375d0dd8", btChannel)
	dichessProfile := rfcomm.NewSerialProfile("dichess", "00001101-0000-1000-8000-00805f9b34fb", btChannel)
	// androidAutoProfile := rfcomm.NewSerialProfile("androidauto", "4de17a00-52cb-11e6-bdf4-0800200c9a66", btChannel)

	if err := dichessProfile.Register(); err != nil {
		log.Fatalf("Could not register android audo profile: %v", err)
	}
	gs.AddShutdownCallback(dichessProfile)
	gs.AddShutdownCallback(server)

	if err := gs.Start(); err != nil {
		log.Fatalf("Could not start graceful shutdown: %v", err)
	}

	if err := bluetooth.StartController(ctx); err != nil {
		log.Fatal(err)
	}

	if hw != nil {
		controller.HardwareGameStarter = hw
		controller.HardwareInput = hw
	}
	if voice != nil {
		controller.VoiceInput = voice
		controller.VoiceGameStarter = voice
	}
	if err := controller.Start(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve())
}
