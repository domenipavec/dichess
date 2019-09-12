package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/Zemanta/gracefulshutdown"
	"github.com/Zemanta/gracefulshutdown/shutdownmanagers/posixsignal"
	wpasupplicant "github.com/dpifke/golang-wpasupplicant"
	"github.com/matematik7/dicar-go/btserver/rfcomm"
	"github.com/matematik7/dichess/go/bluetooth"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/matematik7/dichess/go/hardware"
	"github.com/matematik7/dichess/go/voice"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

var noHardware = flag.Bool("no_hardware", false, "disable hardware init and use fake")

const channel = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()
	observers := &chess_state.Observers{}
	observers.Add(&chess_state.LoggingObserver{})

	hw := hardware.New()
	if !*noHardware {
		if err := hw.Initialize(); err != nil {
			log.Fatal(err)
		}
		observers.Add(hw)
	}

	// for {
	//     data, err := hw.Matrix.Read()
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     for i := range data {
	//         line := ""
	//         for j := range data[i] {
	//             line += chess_state.Square(i, j).String()
	//             line += " "
	//             if data[i][j] {
	//                 line += "+"
	//             } else {
	//                 line += "-"
	//             }
	//             line += " "
	//         }
	//         log.Println(line)
	//     }
	//     log.Println("done")
	//     time.Sleep(time.Second)
	// }

	voice, err := voice.New(ctx)
	if err != nil {
		log.Printf("Couldn't init voice: %v", err)
	} else {
		observers.Add(voice)
	}

	// wpa, err := wpasupplicant.Unixgram("wlan0")
	wpa, err := wpasupplicant.Unixgram("wlx180f76fa4d9a")
	if err != nil {
		log.Fatal(err)
	}

	server := &bluetooth.Server{
		Channel:   channel,
		Observers: observers,
		Wpa:       wpa,
	}

	gs := gracefulshutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	// dichessProfile := rfcomm.NewSerialProfile("dichess", "4f067110-8c71-488a-abf1-5606375d0dd8", channel)
	dichessProfile := rfcomm.NewSerialProfile("dichess", "00001101-0000-1000-8000-00805f9b34fb", channel)
	// androidAutoProfile := rfcomm.NewSerialProfile("androidauto", "4de17a00-52cb-11e6-bdf4-0800200c9a66", channel)

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

	go func() {
		if err := newGame(observers, hw, voice); err != nil {
			log.Println(err)
		}
	}()

	log.Fatal(server.Serve())
}

func newGame(observers *chess_state.Observers, hw *hardware.Hardware, voice *voice.Voice) error {
	player2, err := chess_state.NewUciPlayer()
	if err != nil {
		return err
	}
	player1 := &chess_state.HumanPlayer{
		Inputs: []chess_state.HumanInput{
			hw,
			voice,
		},
	}
	game := chess_state.NewGame(player1, player2, observers)
	go func() {
		for {
			time.Sleep(time.Second)
			data, err := hw.ReadMatrix()
			log.Println(data)
			if err != nil {
				log.Fatal(err)
			}
			done := true
			for i := 0; i < 8; i++ {
				for j := 0; j < 2; j++ {
					if !data[i][j] {
						log.Printf("Missing (%v, %v)", i, j)
						done = false
					}
				}
			}
			for i := 0; i < 8; i++ {
				for j := 6; j < 8; j++ {
					if i == 0 && j == 6 {
						continue
					}
					if i == 2 && j == 6 {
						continue
					}
					if !data[i][j] {
						log.Printf("Missing (%v, %v)", i, j)
						done = false
					}
				}
			}
			if done {
				break
			}
		}
		log.Println("Ready")
		time.Sleep(time.Second)
		if err := voice.Say("What happens now?", texttospeechpb.SsmlVoiceGender_FEMALE); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
		if err := voice.Say("Well, white moves first, and then, we play.", texttospeechpb.SsmlVoiceGender_MALE); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
		if err := game.Play(); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
