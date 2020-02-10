package hardware

import (
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

type Hardware struct {
	initialized bool

	matrix *ReedMatrix
	xAxis  *Axis
	yAxis  *Axis
	coil   *Coil

	movedPieces []movedPiece
}

type movedPiece struct {
	x1, y1 float64
	x2, y2 float64
	color  chess.Color
}

func New() *Hardware {
	return &Hardware{
		matrix: &ReedMatrix{},
		yAxis: &Axis{
			MotorDriver: &MotorDriver{
				Dev: &i2c.Dev{Addr: 21},
			},
			MinusOffset: 220,
			FirstOffset: 0,
			EveryOffset: 247,
			LastOffset:  0,
		},
		xAxis: &Axis{
			MotorDriver: &MotorDriver{
				Dev: &i2c.Dev{Addr: 22},
			},
			MinusOffset: 250,
			FirstOffset: 0,
			EveryOffset: 247,
			LastOffset:  0,
		},
		coil: &Coil{
			MotorDriver: &MotorDriver{
				Dev: &i2c.Dev{Addr: 23},
			},
		},
	}
}

func (h *Hardware) Initialize() error {
	if _, err := host.Init(); err != nil {
		return errors.Wrap(err, "couldn't initialize periph")
	}
	bus, err := i2creg.Open("1")
	if err != nil {
		return errors.Wrap(err, "couldn't open i2c bus 1")
	}
	h.xAxis.MotorDriver.Dev.Bus = bus
	h.yAxis.MotorDriver.Dev.Bus = bus
	h.coil.MotorDriver.Dev.Bus = bus
	h.matrix.Rows = []gpio.PinIn{
		gpioreg.ByName("GPIO4"),
		gpioreg.ByName("GPIO17"),
		gpioreg.ByName("GPIO27"),
		gpioreg.ByName("GPIO22"),
		gpioreg.ByName("GPIO5"),
		gpioreg.ByName("GPIO6"),
		gpioreg.ByName("GPIO13"),
		gpioreg.ByName("GPIO19"),
	}
	h.matrix.Columns = []gpio.PinIO{
		gpioreg.ByName("GPIO21"),
		gpioreg.ByName("GPIO20"),
		gpioreg.ByName("GPIO16"),
		gpioreg.ByName("GPIO12"),
		gpioreg.ByName("GPIO25"),
		gpioreg.ByName("GPIO24"),
		gpioreg.ByName("GPIO23"),
		gpioreg.ByName("GPIO18"),
	}

	// coil should be initialized first so it doesn't heat for nothing
	if err := h.coil.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize coil")
	}

	if err := h.matrix.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize reed matrix")
	}

	// Must home axis with addr 21 before addr 22
	if err := h.yAxis.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize y-axis")
	}
	if err := h.xAxis.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize x-axis")
	}

	h.initialized = true

	return nil
}
