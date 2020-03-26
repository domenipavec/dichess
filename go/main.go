package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/Zemanta/gracefulshutdown"
	"github.com/Zemanta/gracefulshutdown/shutdownmanagers/posixsignal"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/matematik7/dichess/go/bluetooth"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/matematik7/dichess/go/hardware"
	"github.com/matematik7/dichess/go/voice"
	"github.com/matematik7/dichess/go/wpa"
	"github.com/tj/go-update"
	"github.com/tj/go-update/stores/github"
)

const currentVersion = "0.1.3"

var (
	noHardware = flag.Bool("no_hardware", false, "disable hardware init and use fake")
	noUpdate   = flag.Bool("no_update", false, "disable updates")
	cpuprofile = flag.Bool("cpuprofile", false, "write cpu profile to file")
)

const btChannel = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	if !*noUpdate {
		go func() {
			if err := doUpdate(currentVersion); err != nil {
				log.Println(err)
			}
		}()
	}

	if *cpuprofile {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
	}

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

	stateSenders := &chess_state.StateSenders{}
	stateSenders.Add(&chess_state.LoggingStateSender{})

	controller := &chess_state.Controller{Observers: observers, StateSenders: stateSenders}
	server := bluetooth.NewServer(btChannel, controller, wpa)
	controller.BluetoothInput = server

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
	gs.AddShutdownCallback(gracefulshutdown.ShutdownFunc(func(string) error { pprof.StopCPUProfile(); return nil }))

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

// Installs the new binary, will run on next restart
func doUpdate(current string) error {
	log.Printf("Current version %s", current)
	for {
		manager := &update.Manager{
			Store: &github.Store{
				Owner:   "matematik7",
				Repo:    "dichess",
				Version: current,
			},
			Command: "./dichess",
		}

		latest, err := manager.LatestReleases()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		if len(latest) == 0 {
			log.Println("No updates.")
			return nil
		}

		log.Printf("Updating to %s", latest[0].Version)

		asset := latest[0].FindTarball(runtime.GOOS, runtime.GOARCH)
		if asset == nil {
			log.Println("No binary for your system.")
			return nil
		}

		path, err := asset.Download()
		if err != nil {
			return err
		}

		if err := manager.Install(path); err != nil {
			return err
		}

		log.Printf("Updated to %s", latest[0].Version)

		return nil
	}

	return nil
}
