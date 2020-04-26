package hardware

import (
	"log"
	"time"

	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

type Initializer interface {
	Initialize() error
}

type Matrix interface {
	Initializer
	Read() ([][]bool, error)
}

type Axis interface {
	Initializer
	SetCurrent(uint8) error
	GoTo(float64, uint8) error
}

type Coil interface {
	Initializer
	On() error
	Off() error
	SetPwm(uint8) error
	Rotate(int) error
}

type Hardware struct {
	Settings bluetoothpb.SettingsProvider

	initialized bool
	fake        bool

	matrix Matrix
	xAxis  Axis
	yAxis  Axis
	coil   Coil

	movedPieces []movedPiece
}

type movedPiece struct {
	x1, y1 float64
	x2, y2 float64
	piece  chess.Piece
}

func New() *Hardware {
	return &Hardware{}
}

func (h *Hardware) Initialize() error {
	for _, initializer := range []Initializer{h.matrix, h.coil, h.yAxis, h.xAxis} {
		if err := initializer.Initialize(); err != nil {
			return err
		}
	}
	h.initialized = true
	return nil
}

func (h *Hardware) InitializeReal() error {
	if _, err := host.Init(); err != nil {
		return errors.Wrap(err, "couldn't initialize periph")
	}
	bus, err := i2creg.Open("1")
	if err != nil {
		return errors.Wrap(err, "couldn't open i2c bus 1")
	}
	h.yAxis = &RealAxis{
		MotorDriver: &MotorDriver{
			Dev: &i2c.Dev{Addr: 21, Bus: bus},
		},
		MinusOffset: 195,
		FirstOffset: 0,
		EveryOffset: 250,
		LastOffset:  0,
	}
	h.xAxis = &RealAxis{
		MotorDriver: &MotorDriver{
			Dev: &i2c.Dev{Addr: 22, Bus: bus},
		},
		MinusOffset: 225,
		FirstOffset: 0,
		EveryOffset: 250,
		LastOffset:  0,
	}
	h.coil = &RealCoil{
		MotorDriver: &MotorDriver{
			Dev: &i2c.Dev{Addr: 23, Bus: bus},
		},
	}

	h.matrix = &ReedMatrix{
		Rows: []gpio.PinIn{
			gpioreg.ByName("GPIO4"),
			gpioreg.ByName("GPIO17"),
			gpioreg.ByName("GPIO27"),
			gpioreg.ByName("GPIO22"),
			gpioreg.ByName("GPIO5"),
			gpioreg.ByName("GPIO6"),
			gpioreg.ByName("GPIO13"),
			gpioreg.ByName("GPIO19"),
		},
		Columns: []gpio.PinIO{
			gpioreg.ByName("GPIO0"),
			gpioreg.ByName("GPIO26"),
			gpioreg.ByName("GPIO16"),
			gpioreg.ByName("GPIO12"),
			gpioreg.ByName("GPIO25"),
			gpioreg.ByName("GPIO24"),
			gpioreg.ByName("GPIO23"),
			gpioreg.ByName("GPIO18"),
		},
	}

	return h.Initialize()
}

func (h *Hardware) InitializeFake() error {
	h.xAxis = &FakeAxis{}
	h.yAxis = &FakeAxis{}
	h.coil = &FakeCoil{}
	h.matrix = &FakeMatrix{}
	h.fake = true

	return h.Initialize()
}

func (h *Hardware) SensorTest() {
	for {
		data, err := h.matrix.Read()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Sensor data:")
		for _, line := range data {
			log.Println(line)
		}
		time.Sleep(time.Second)
	}
}

// func (h *Hardware) Test() {
//     h.coil.SetPwm(100)
//     defer h.coil.Off()
//     defer h.xAxis.SetCurrent(10)
//     defer h.yAxis.SetCurrent(10)
//
//     h.xAxis.SetCurrent(77)
//     h.yAxis.SetCurrent(77)
//     h.Do(
//         func() error { return h.xAxis.GoTo(0, 40) },
//         func() error { return h.yAxis.GoTo(0, 40) },
//     )
//     h.xAxis.SetCurrent(10)
//
//     for {
//         for i := 0; i < 8; i++ {
//             h.coil.On()
//             h.yAxis.GoTo(float64(i), 40)
//             h.coil.Off()
//             time.Sleep(time.Second)
//             h.coil.On()
//             time.Sleep(time.Second)
//             h.coil.Off()
//         }
//         for i := 7; i >= 0; i-- {
//             h.coil.On()
//             h.yAxis.GoTo(float64(i), 40)
//             h.coil.Off()
//             time.Sleep(time.Second)
//             h.coil.On()
//             time.Sleep(time.Second)
//             h.coil.Off()
//         }
//     }
// }
