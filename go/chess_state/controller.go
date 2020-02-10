package chess_state

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/pkg/errors"
)

type GameStarter interface {
	StartGame() error
}

type Controller struct {
	HardwareInput HumanInput
	VoiceInput    HumanInput

	HardwareObserver Observer
	VoiceObserver    Observer

	HardwareGameStarter GameStarter
	VoiceGameStarter    GameStarter

	settingsMutex sync.Mutex
	settings      *bluetoothpb.Settings

	game *Game
}

func (c *Controller) GetSettings() *bluetoothpb.Settings {
	c.settingsMutex.Lock()
	defer c.settingsMutex.Unlock()

	return c.settings
}

func (c *Controller) SetSettings(s *bluetoothpb.Settings) {
	c.settingsMutex.Lock()
	defer c.settingsMutex.Unlock()

	c.settings = s

	if err := saveSettings(s); err != nil {
		log.Printf("Could not save settings: %v", err)
	}
}

func (c *Controller) GetGame() *Game {
	return c.game
}

func (c *Controller) Start() error {
	settings, err := loadSettings()
	if err != nil {
		return err
	}
	c.settings = settings

	go c.gameStarter()

	return nil
}

func (c *Controller) gameStarter() {
	for range time.NewTicker(time.Second).C {
		if c.GetGame() != nil {
			continue
		}
		if err := c.StartGame(); err != nil {
			log.Println(err)
		}
	}
}

func (c *Controller) createPlayer(s *bluetoothpb.Settings, typ bluetoothpb.Settings_PlayerType) (Player, error) {
	switch typ {
	case bluetoothpb.Settings_HUMAN:
		inputs := []HumanInput{}
		if c.HardwareInput != nil {
			inputs = append(inputs, c.HardwareInput)
		}
		if s.VoiceRecognition && c.VoiceInput != nil {
			inputs = append(inputs, c.VoiceInput)
		}
		return &HumanPlayer{
			Inputs: inputs,
		}, nil
	case bluetoothpb.Settings_COMPUTER:
		return NewUciPlayer(s.ComputerSettings)
	default:
		return nil, errors.Errorf("invalid player type: %v", typ)
	}

}

func (c *Controller) StartGame() error {
	if c.GetGame() != nil {
		return errors.New("game in progress")
	}

	settings := c.GetSettings()

	observers := &Observers{}
	observers.Add(&LoggingObserver{})

	player1, err := c.createPlayer(settings, settings.Player1)
	if err != nil {
		return errors.Wrap(err, "could not create player 1")
	}
	player2, err := c.createPlayer(settings, settings.Player2)
	if err != nil {
		return errors.Wrap(err, "could not create player 2")
	}

	if settings.RandomBw && rand.Intn(2) == 0 {
		player1, player2 = player2, player1
	}

	if settings.AutoMove && c.HardwareObserver != nil {
		observers.Add(c.HardwareObserver)
	}
	if settings.Sound && c.VoiceObserver != nil {
		observers.Add(c.VoiceObserver)
	}

	if err := c.HardwareGameStarter.StartGame(); err != nil {
		return err
	}
	if settings.Sound {
		if err := c.VoiceGameStarter.StartGame(); err != nil {
			return err
		}
	}

	game := NewGame(player1, player2, observers)
	go func() {
		if err := game.Play(); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
