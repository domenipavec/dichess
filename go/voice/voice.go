package voice

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	speech "cloud.google.com/go/speech/apiv1"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

var moveString = map[string]string{
	"sl-SI": "%s na %s",
	"en-US": "%s to %s",
}
var piecesStrings = map[string]map[chess.Piece][]string{
	"sl-SI": map[chess.Piece][]string{
		chess.WhiteKing:   []string{"kralj"},
		chess.WhiteQueen:  []string{"kraljica", "dama"},
		chess.WhiteRook:   []string{"top", "trdnjava"},
		chess.WhiteBishop: []string{"lovec", "tekač", "laufer"},
		chess.WhiteKnight: []string{"skakač", "konj"},
		chess.WhitePawn:   []string{"kmet"},
	},
	"en-US": map[chess.Piece][]string{
		chess.WhiteKing:   []string{"king"},
		chess.WhiteQueen:  []string{"queen"},
		chess.WhiteRook:   []string{"rook", "castle"},
		chess.WhiteBishop: []string{"bishop"},
		chess.WhiteKnight: []string{"knight"},
		chess.WhitePawn:   []string{"pawn"},
	},
}

type Voice struct {
	speechClient *speech.Client
	ttsClient    *texttospeech.Client
	Language     string
}

func New(ctx context.Context) (*Voice, error) {
	speechClient, err := speech.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't initialize speech client")
	}
	ttsClient, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't initialize text-to-speech client")
	}

	return &Voice{
		speechClient: speechClient,
		ttsClient:    ttsClient,

		// Language: "sl-SI",
		Language: "en-US",
	}, nil
}

func (v *Voice) Update(game *chess.Game, move *chess_state.Move) error {
	if len(game.Moves()) < 1 {
		return nil
	}
	if !move.ShouldSay {
		return nil
	}
	gameMove := game.Moves()[len(game.Moves())-1]
	piece := game.Position().Board().Piece(gameMove.S2())
	if piece == chess.NoPiece {
		log.Printf("Wierd move with no piece: %v", move)
		return nil
	}
	// normalize to white pieces for getting string
	if piece > chess.WhitePawn {
		piece -= chess.WhitePawn
	}

	txt := fmt.Sprintf(moveString[v.Language], piecesStrings[v.Language][piece], gameMove.S2().String())
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

func (v *Voice) Say(txt string, gender texttospeechpb.SsmlVoiceGender) error {
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
	defer log.Println("recognize done")
	for {
		log.Println("recognize loop")
		ctx := ctx
		select {
		case <-ctx.Done():
			log.Println("recognize ctx done")
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
			defer log.Println("recv done")
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
					log.Println(result)
					for _, alternative := range result.Alternatives {
						for _, phrase := range phrases {
							if strings.ToLower(alternative.Transcript) == strings.ToLower(phrase) {
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

		log.Println(phrases)
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
		log.Printf("Result: %v", result)
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

	log.Println("send start")
	defer log.Println("send done")
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
