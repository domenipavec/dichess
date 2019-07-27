package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func voiceRecognition(ctx context.Context) error {
	client, err := speech.NewClient(ctx)
	if err != nil {
		return err
	}

	for {

		recognize, err := client.StreamingRecognize(ctx)
		if err != nil {
			return err
		}

		// This context is not for request, but for canceling send loop
		ctx, cancel := context.WithCancel(ctx)

		go func() {
			defer cancel()
			for {
				response, err := recognize.Recv()
				if err != nil {
					log.Printf("Receive err: %v", err)
					return
				}
				if response.GetSpeechEventType() == speechpb.StreamingRecognizeResponse_END_OF_SINGLE_UTTERANCE {
					log.Println("Closing done")
					cancel()
				}
				log.Printf("results: %v", len(response.Results))
				for _, result := range response.Results {
					log.Printf("final: %v", result.GetIsFinal())
					for _, alternative := range result.Alternatives {
						log.Printf("alternative (%v): %v", alternative.Confidence, alternative.Transcript)
					}
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
						// LanguageCode:    "en-US",
						// SpeechContexts: []*speechpb.SpeechContext{
						//     &speechpb.SpeechContext{
						//         Phrases: []string{
						//             "knight to",
						//             "pawn to",
						//             "rook to",
						//             "bishop to",
						//             "queen to",
						//             "king to",
						//             "A1",
						//             "A2",
						//             "A3",
						//             "B1",
						//             "B2",
						//             "B3",
						//             "C1",
						//             "C2",
						//             "C3",
						//             "D1",
						//             "D2",
						//             "D3",
						//             "E1",
						//             "E2",
						//             "E3",
						//             "F1",
						//             "F2",
						//             "F3",
						//         },
						//     },
						// },
						LanguageCode: "sl-SI",
						SpeechContexts: []*speechpb.SpeechContext{
							&speechpb.SpeechContext{
								Phrases: []string{
									"konj na",
									"skakač na",
									"trdnjava na",
									"top na",
									"kralj na",
									"kraljica na",
									"kmet na",
									"A1",
									"A2",
									"A3",
									"B1",
									"B2",
									"B3",
									"C1",
									"C2",
									"C3",
									"D1",
									"D2",
									"D3",
									"E1",
									"E2",
									"E3",
									"F1",
									"F2",
									"F3",
								},
							},
						},
						Model: "command_and_search",
					},
					SingleUtterance: true,
				},
			},
		}); err != nil {
			return errors.Wrap(err, "could not send recognition config")
		}

		if err := voiceSendLoop(ctx, recognize); err != nil {
			return err
		}
	}

	return nil
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
	go cmd.Wait()

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
	return nil
}
