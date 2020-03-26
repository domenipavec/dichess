package voice

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func (v *Voice) Update(_ context.Context, stateSender chess_state.StateSender, game *chess_state.Game, move *chess_state.Move) error {
	if len(game.Game.Moves()) < 1 {
		return nil
	}
	if !move.ShouldSay {
		return nil
	}
	if move.Undo {
		// TODO: say undo?
		return nil
	}
	gameMove := game.Game.Moves()[len(game.Game.Moves())-1]
	piece := game.Game.Position().Board().Piece(gameMove.S2())
	if piece == chess.NoPiece {
		log.Printf("Wierd move with no piece: %v", move)
		return nil
	}
	// normalize to white pieces for getting string
	if piece > chess.WhitePawn {
		piece -= chess.WhitePawn
	}

	txt := fmt.Sprintf("%s %s %s", piecesStrings[v.Language][piece][0], translations[v.Language]["to"], gameMove.S2().String())
	if err := v.Say(txt, texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
		return err
	}
	if gameMove.HasTag(chess.Check) {
		if err := v.Say("check!", texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
			return err
		}
	}
	return nil
}

func (v *Voice) StartGame(stateSender chess_state.StateSender) error {
	stateSender.StateSend("Game starting.")

	time.Sleep(time.Second)
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
			LanguageCode: v.Language,
			SsmlGender:   gender,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
			SpeakingRate:  0.5,
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
