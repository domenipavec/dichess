package voice

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func (v *Voice) Update(_ context.Context, stateSender chess_state.StateSender, game *chess_state.Game, move *chess_state.Move) error {
	if len(game.Game.Moves()) < 1 {
		if v.Settings.GetSettings().Intro {
			if err := v.intro(stateSender); err != nil {
				return err
			}
		}
		return nil
	}
	if move == nil {
		return nil
	}
	if move.Undo {
		if err := v.Say("Move undo, please move manually.", texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
			return err
		}
		return nil
	}
	gameMove := game.Game.Moves()[len(game.Game.Moves())-1]
	if gameMove.HasTag(chess.KingSideCastle) || gameMove.HasTag(chess.QueenSideCastle) {
		if err := v.Say("Castling, move manually please.", texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
			return err
		}
	}
	if gameMove.HasTag(chess.Capture) {
		if err := v.Say(fmt.Sprintf("Piece capture moving to %s, please remove.", gameMove.S2().String()), texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
			return err
		}
	}
	if !move.ShouldSay {
		return nil
	}
	piece := game.Game.Position().Board().Piece(gameMove.S2())
	if piece == chess.NoPiece {
		log.Printf("Wierd move with no piece: %v", move)
		return nil
	}
	// normalize to white pieces for getting string
	if piece > chess.WhitePawn {
		piece -= chess.WhitePawn
	}

	lang := bluetoothpb.Settings_ENGLISH
	txt := fmt.Sprintf("%s %s %s", piecesStrings[lang][piece][0], translations[lang]["to"], gameMove.S2().String())
	if err := v.Say(txt, texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
		return err
	}
	if game.Game.Outcome() != chess.NoOutcome {
		if err := v.Say(game.Game.Method().String(), texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
			return err
		}
	} else if gameMove.HasTag(chess.Check) {
		if err := v.Say("check!", texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
			return err
		}
	}
	return nil
}

func (v *Voice) intro(stateSender chess_state.StateSender) error {
	stateSender.StateSend("Game starting.")

	time.Sleep(time.Second)
	stateSender.StateSend("Game starting.")
	if err := v.Say("What happens now?", texttospeechpb.SsmlVoiceGender_FEMALE); err != nil {
		return err
	}
	time.Sleep(time.Second)
	if err := v.Say("Well, white moves first, and then, we play.", texttospeechpb.SsmlVoiceGender_MALE); err != nil {
		return err
	}
	time.Sleep(time.Second)

	return nil
}

func (v *Voice) Say(txt string, gender texttospeechpb.SsmlVoiceGender) error {
	log.Println("Saying: ", txt)
	if !v.Settings.GetSettings().Sound {
		return nil
	}
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: txt},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   gender,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
			SpeakingRate:  0.7,
		},
	}

	resp, err := v.ttsClient.SynthesizeSpeech(context.Background(), &req)
	if err != nil {
		return errors.Wrap(err, "couldn't synthesize speech")
	}

	cmd := exec.Command("aplay", "-")
	cmd.Stdin = bytes.NewReader(resp.AudioContent)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "could not run aplay")
	}

	return nil
}
