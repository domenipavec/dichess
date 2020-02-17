package chess_state

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/pkg/errors"
)

func loadSettings() (*bluetoothpb.Settings, error) {
	s := &bluetoothpb.Settings{
		Sound:            true,
		Language:         bluetoothpb.Settings_ENGLISH,
		VoiceRecognition: true,
		AutoMove:         true,
		RandomBw:         true,
		Player1:          bluetoothpb.Settings_COMPUTER,
		Player2:          bluetoothpb.Settings_HUMAN,
		ComputerSettings: &bluetoothpb.Settings_ComputerSettings{
			TimeLimitMs:   1000,
			SkillLevel:    20,
			LimitStrength: false,
			Elo:           1350,
		},
	}
	data, err := ioutil.ReadFile("settings.proto")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return s, saveSettings(s)
		}
		return nil, errors.Wrap(err, "could not read settings.proto")
	}
	if err := proto.UnmarshalText(string(data), s); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal settings")
	}
	return s, nil
}

func saveSettings(s *bluetoothpb.Settings) error {
	f, err := os.Create("settings.proto")
	if err != nil {
		return errors.Wrap(err, "could not create settings.proto")
	}
	defer f.Close()

	if err := proto.MarshalText(f, s); err != nil {
		return errors.Wrap(err, "could not marshal settings to text")
	}
	return nil
}
