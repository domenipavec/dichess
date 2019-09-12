package hardware

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/pkg/errors"
	"periph.io/x/periph/conn/i2c"
)

const (
	PinInverted = 1 << iota
	PinPullUp
	PinOutput
	PinHigh
	PinEnabled
	PinHome
	PinNotUsed
	Pin2
)

type MotorDriver struct {
	*i2c.Dev
}

func (d *MotorDriver) Home() error {
	if err := d.Tx([]byte{0x01}, nil); err != nil {
		return errors.Wrap(err, "couldn't initialize homing")
	}
	return d.Wait()
}

func (d *MotorDriver) Wait() error {
	status := []byte{0x00}
	for status[0]&0x01 != 0x01 {
		time.Sleep(1000 * time.Millisecond)
		if err := d.Tx(nil, status); err != nil {
			return errors.Wrap(err, "couldn't read status")
		}
	}
	return nil
}

func (d *MotorDriver) GetPosition() (int16, error) {
	data := make([]byte, 2)
	if err := d.Tx([]byte{1}, data); err != nil {
		return 0, errors.Wrap(err, "couldn't read position")
	}

	var value int16
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &value); err != nil {
		return 0, errors.Wrap(err, "couldn't decode position")
	}

	return value, nil
}

// SetSpeed value is inverse of actual speed, 90 is normal, 10 is really fast
func (d *MotorDriver) SetSpeed(speed uint8) error {
	if err := d.Tx([]byte{0x02, speed}, nil); err != nil {
		return errors.Wrap(err, "couldn't set speed")
	}
	return nil
}

// current is 0.01176A*value
func (d *MotorDriver) SetCurrent(value uint8) error {
	if err := d.Tx([]byte{0x04, value}, nil); err != nil {
		return errors.Wrap(err, "couldn't set current")
	}
	time.Sleep(time.Millisecond * 300)
	return nil
}

func (d *MotorDriver) SetPin(flags uint8) error {
	if err := d.Tx([]byte{0x05, flags}, nil); err != nil {
		return errors.Wrap(err, "couldn't set flags")
	}
	return nil
}

func (d *MotorDriver) Go(dest int16) error {
	var buf bytes.Buffer
	// write command, cannot fail
	buf.Write([]byte{0x03})
	if err := binary.Write(&buf, binary.LittleEndian, dest); err != nil {
		return errors.Wrap(err, "couldn't encode destination")
	}
	if err := d.Tx(buf.Bytes(), nil); err != nil {
		return errors.Wrap(err, "couldn't set destination")
	}
	return d.Wait()
}

type Coil struct {
	*MotorDriver
}

func (c *Coil) Initialize() error {
	if err := c.SetPin(PinOutput); err != nil {
		return errors.Wrap(err, "couldn't set pin 1 as output")
	}
	if err := c.SetPin(Pin2 | PinPullUp | PinEnabled | PinHome); err != nil {
		return errors.Wrap(err, "couldn't set pin 2 flags")
	}
	if err := c.SetSpeed(80); err != nil {
		return errors.Wrap(err, "couldn't set speed for coil")
	}
	if err := c.Home(); err != nil {
		return err
	}
	// Disable homing pin after homing
	if err := c.SetPin(Pin2 | PinPullUp); err != nil {
		return errors.Wrap(err, "couldn't set pin 2 flags")
	}
	if err := c.Rotate(0); err != nil {
		return err
	}
	return nil
}

func (c *Coil) On() error {
	return errors.Wrap(c.SetPin(PinOutput|PinHigh), "couldn't turn coil on")
}

func (c *Coil) Off() error {
	return errors.Wrap(c.SetPin(PinOutput), "couldn't turn coil on")
}

func (c *Coil) Rotate(degrees int) error {
	return errors.Wrap(c.Go(int16(degrees)*950/90+950), "couldn't rotate coil")
}

type Axis struct {
	*MotorDriver

	MinusOffset int
	FirstOffset int
	EveryOffset int
	LastOffset  int
}

func (a *Axis) Initialize() error {
	// Axis are running old code, otherwise these should be inverted
	if err := a.SetPin(PinPullUp | PinEnabled | PinHome); err != nil {
		return errors.Wrap(err, "couldn't set pin 1 flags")
	}
	if err := a.SetPin(Pin2 | PinPullUp | PinEnabled | PinHome); err != nil {
		return errors.Wrap(err, "couldn't set pin 1 flags")
	}
	return a.Home()
}

func (a *Axis) Home() error {
	if err := a.SetCurrent(50); err != nil {
		return errors.Wrap(err, "couldn't set moving current")
	}
	if err := a.SetSpeed(80); err != nil {
		return errors.Wrap(err, "couldn't set moving speed")
	}
	if err := a.MotorDriver.Home(); err != nil {
		return err
	}

	for {
		pos, err := a.GetPosition()
		if err != nil {
			return err
		}
		if pos == 20 {
			break
		}
		if err := a.Go(20); err != nil {
			return err
		}
	}

	if err := a.SetCurrent(10); err != nil {
		return errors.Wrap(err, "couldn't set idle current")
	}
	return nil
}

func (a *Axis) getDest(id float64) int {
	if id < -1 {
		return a.MinusOffset
	}
	if id < 0 {
		return a.MinusOffset + int((id+1)*float64(a.FirstOffset))
	}
	if id <= 7 {
		return a.MinusOffset + a.FirstOffset + int(id*float64(a.EveryOffset))
	}
	if id < 8 {
		return a.MinusOffset + a.FirstOffset + 7*a.EveryOffset + int((id-7)*float64(a.LastOffset))
	}
	return a.MinusOffset + a.FirstOffset + 7*a.EveryOffset + a.LastOffset
}

func (a *Axis) GoTo(id float64, speed uint8) error {
	if err := a.SetSpeed(speed); err != nil {
		return errors.Wrap(err, "couldn't set moving speed")
	}
	if err := a.Go(int16(a.getDest(id))); err != nil {
		return err
	}
	return nil
}
