package voice

import (
	"context"

	speech "cloud.google.com/go/speech/apiv1"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/pkg/errors"
)

type Voice struct {
	Settings bluetoothpb.SettingsProvider

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
