package voice

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func (v *Voice) MakeMove(ctx context.Context, stateSender chess_state.StateSender, game *chess.Game) (*chess_state.Move, error) {
	if !v.Settings.GetSettings().VoiceRecognition {
		<-ctx.Done()
		return nil, ctx.Err()
	}
	lang := v.Settings.GetSettings().Language
	phrases := v.generatePhrases(game, lang)

	for {
		result, err := v.recognize(ctx, phrases, languages[lang])
		if err != nil {
			return nil, err
		}
		log.Printf("result: %v", result)

		moves, err := v.parseMove(game, result, lang)
		if err != nil {
			return nil, err
		}
		log.Printf("moves: %v", moves)

		if len(moves) < 1 {
			stateSender.StateSend(fmt.Sprintf("Could not recognize move: %s", result))
			continue
		}
		if len(moves) > 1 {
			construct := ""
			for i := range moves {
				if i > 0 {
					if i < len(moves)-2 {
						construct += ", "
					} else {
						construct += " or "
					}
				}
				construct += moves[i].S1().String()
			}
			prompt := fmt.Sprintf("%s "+translations[bluetoothpb.Settings_ENGLISH]["ambiguous"]+" %s?", result, construct)
			stateSender.StateSend(prompt)

			if err := v.Say(prompt, texttospeechpb.SsmlVoiceGender_NEUTRAL); err != nil {
				return nil, err
			}

			fields := make([]string, 2*len(moves))
			for i, move := range moves {
				fields[i] = move.S1().String()
				fields[i] = translations[lang]["from"] + " " + move.S1().String()
			}
			from, err := v.recognize(ctx, fields, languages[lang])
			if err != nil {
				return nil, err
			}

			found := false
			for _, move := range moves {
				if strings.Contains(from, move.S1().String()) {
					found = true
					moves[0] = move
				}
			}
			if !found {
				stateSender.StateSend(fmt.Sprintf("Could not recognize move: %s from %s", result, from))
				continue
			}
		}

		return &chess_state.Move{
			Move:       moves[0],
			ShouldMove: true,
			ShouldSay:  true,
		}, nil
	}
}

func (v *Voice) recognize(ctx context.Context, phrases []string, language string) (string, error) {
	resultChan := make(chan string, 1)
	for {
		ctx := ctx
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		ctx, cancelTimeout := context.WithTimeout(ctx, time.Second*15)
		defer cancelTimeout()

		recognize, err := v.speechClient.StreamingRecognize(ctx)
		if err != nil {
			return "", err
		}

		// This context is not for request, but for canceling send loop
		ctx, cancel := context.WithCancel(ctx)

		go func() {
			defer cancel()
			for {
				response, err := recognize.Recv()
				if err != nil {
					log.Printf("Receive err: %v", err)
					resultChan <- ""
					return
				}
				for _, result := range response.Results {
					if !result.GetIsFinal() {
						continue
					}
					for _, alternative := range result.Alternatives {
						for _, phrase := range phrases {
							if strings.EqualFold(alternative.Transcript, phrase) {
								resultChan <- phrase
								return
							}
						}
					}
					resultChan <- ""
					return
				}
			}
		}()

		log.Printf("Phrases: %v", phrases)
		if err := recognize.Send(&speechpb.StreamingRecognizeRequest{
			StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
				StreamingConfig: &speechpb.StreamingRecognitionConfig{
					Config: &speechpb.RecognitionConfig{
						Encoding:        speechpb.RecognitionConfig_LINEAR16,
						SampleRateHertz: 44100,
						MaxAlternatives: 5,
						LanguageCode:    language,
						SpeechContexts: []*speechpb.SpeechContext{
							&speechpb.SpeechContext{
								Phrases: phrases,
							},
						},
						Model: "command_and_search",
					},
					SingleUtterance: true,
				},
			},
		}); err != nil {
			return "", errors.Wrap(err, "could not send recognition config")
		}

		if err := voiceSendLoop(ctx, recognize); err != nil {
			return "", err
		}

		result := <-resultChan
		if result != "" {
			return result, nil
		}
	}
}

func voiceSendLoop(ctx context.Context, recognize speechpb.Speech_StreamingRecognizeClient) error {
	cmd := exec.CommandContext(ctx, "arecord", "-r", "44100", "-c", "1", "-f", "S16_LE", "-t", "wav", "--device=hw:1,0", "-")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "could not create stdout pipe")
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "could not start arecord")
	}
	defer func() {
		if err := stdout.Close(); err != nil {
			log.Printf("Closing of pipe err: %v", err)
		}
	}()
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("arecord exited with: %v", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		data := make([]byte, 1024)
		n, err := stdout.Read(data)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			if strings.Contains(err.Error(), "file already closed") {
				return nil
			}
			return errors.Wrap(err, "stdout read")
		}
		data = data[:n]

		if err := recognize.Send(&speechpb.StreamingRecognizeRequest{
			StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
				AudioContent: data,
			},
		}); err != nil {
			return errors.Wrap(err, "recognize send")
		}
	}
}

func (v *Voice) generatePhrases(game *chess.Game, lang bluetoothpb.Settings_Language) []string {
	phraseMap := make(map[string]bool)

	for _, move := range game.ValidMoves() {
		piece := game.Position().Board().Piece(move.S1())
		if piece == chess.NoPiece {
			log.Printf("Wierd move with no piece: %v", move)
			continue
		}
		// normalize to white pieces for getting string
		if piece > chess.WhitePawn {
			piece -= chess.WhitePawn
		}
		for _, pieceStr := range piecesStrings[lang][piece] {
			current := fmt.Sprintf("%s %s %s", pieceStr, translations[lang]["to"], move.S2().String())
			phraseMap[current] = true
		}
	}

	phrases := make([]string, 0, len(phraseMap))
	for phrase := range phraseMap {
		phrases = append(phrases, phrase)
	}
	return phrases
}

func (v *Voice) parseMove(game *chess.Game, phrase string, lang bluetoothpb.Settings_Language) ([]*chess.Move, error) {
	var pieceStr, squareStr string
	if _, err := fmt.Sscanf(phrase, "%s "+translations[lang]["to"]+" %s", &pieceStr, &squareStr); err != nil {
		return nil, errors.Wrapf(err, "couldn't parse phrase %v", phrase)
	}

	piece := chess.NoPiece
	for searchPiece, searchStrings := range piecesStrings[lang] {
		for _, searchString := range searchStrings {
			if searchString == pieceStr {
				piece = searchPiece
				break
			}
		}
		if piece != chess.NoPiece {
			break
		}
	}
	if piece == chess.NoPiece {
		return nil, errors.Errorf("couldn't find piece %s", pieceStr)
	}
	if game.Position().Turn() == chess.Black {
		piece += chess.WhitePawn
	}

	var moves []*chess.Move
	for _, move := range game.ValidMoves() {
		if piece == game.Position().Board().Piece(move.S1()) && move.S2().String() == squareStr {
			moves = append(moves, move)
		}
	}
	return moves, nil
}
