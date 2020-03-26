package chess_state

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/pkg/errors"
)

type GameStarter interface {
	StartGame(StateSender) error
}

type Controller struct {
	Inputs []HumanInput

	HardwareGameStarter GameStarter
	VoiceGameStarter    GameStarter

	Observers    *Observers
	StateSenders *StateSenders

	settingsMutex sync.Mutex
	settings      *bluetoothpb.Settings

	game   *Game
	cancel func()
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
		return &HumanPlayer{
			Inputs: c.Inputs,
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

	if err := c.HardwareGameStarter.StartGame(c.StateSenders); err != nil {
		return err
	}

	settings := c.GetSettings()

	if c.VoiceGameStarter != nil && settings.Sound {
		if err := c.VoiceGameStarter.StartGame(c.StateSenders); err != nil {
			return err
		}
	}

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

	// if settings.AutoMove && c.HardwareObserver != nil {
	//     observers.Add(c.HardwareObserver)
	// }
	// if settings.Sound && c.VoiceObserver != nil {
	//     observers.Add(c.VoiceObserver)
	// }

	c.game = NewGame(player1, player2, c.Observers, c.StateSenders)
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel
	go func() {
		defer cancel()
		defer func() {
			c.game = nil
		}()
		if err := c.game.Play(ctx); err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (c *Controller) StopGame() {
	if c.cancel != nil {
		c.cancel()
	}
}
