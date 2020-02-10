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

	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func (v *Voice) MakeMove(ctx context.Context, game *chess.Game) (*chess_state.Move, error) {
	phrases := v.generatePhrases(game)
	result, err := v.recognize(ctx, phrases)
	if err != nil {
		return nil, err
	}

	moves, err := v.parseMove(game, result)
	if err != nil {
		return nil, err
	}

	if len(moves) < 1 {
		log.Println("Invalid move in speech")
		return nil, nil
	}
	if len(moves) > 1 {
		// TODO disambiguate
		log.Println("Voice ambiguous move")
		return nil, nil
	}

	if err := game.Move(moves[0]); err != nil {
		return nil, errors.Wrap(err, "couldn't make voice move")
	}

	return &chess_state.Move{
		ShouldMove: true,
		ShouldSay:  false,
	}, nil
}

func (v *Voice) recognize(ctx context.Context, phrases []string) (string, error) {
	resultChan := make(chan string, 1)
	for {
		ctx := ctx
		select {
		case <-ctx.Done():
			return "", nil
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

		if err := recognize.Send(&speechpb.StreamingRecognizeRequest{
			StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
				StreamingConfig: &speechpb.StreamingRecognitionConfig{
					Config: &speechpb.RecognitionConfig{
						Encoding:        speechpb.RecognitionConfig_LINEAR16,
						SampleRateHertz: 44100,
						MaxAlternatives: 5,
						LanguageCode:    v.Language,
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
			return nil
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

func (v *Voice) generatePhrases(game *chess.Game) []string {
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
		for _, pieceStr := range piecesStrings[v.Language][piece] {
			current := fmt.Sprintf(moveString[v.Language], pieceStr, move.S2().String())
			phraseMap[current] = true
		}
	}

	phrases := make([]string, 0, len(phraseMap))
	for phrase := range phraseMap {
		phrases = append(phrases, phrase)
	}
	return phrases
}

func (v *Voice) parseMove(game *chess.Game, phrase string) ([]*chess.Move, error) {
	var pieceStr, squareStr string
	if _, err := fmt.Sscanf(phrase, moveString[v.Language], &pieceStr, &squareStr); err != nil {
		return nil, errors.Wrapf(err, "couldn't parse phrase %v", phrase)
	}

	piece := chess.NoPiece
	for searchPiece, searchStrings := range piecesStrings[v.Language] {
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
		return nil, errors.Errorf("couldn't find piece %v", pieceStr)
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
