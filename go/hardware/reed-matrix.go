package hardware

import (
	"time"

	"github.com/pkg/errors"
	"periph.io/x/periph/conn/gpio"
)

type ReedMatrix struct {
	Columns []gpio.PinIO
	Rows    []gpio.PinIn
}

func (m *ReedMatrix) Initialize() error {
	for i, col := range m.Columns {
		if col == nil {
			return errors.Errorf("invalid gpio for column %d", i)
		}
		if err := col.In(gpio.Float, gpio.NoEdge); err != nil {
			return errors.Wrapf(err, "couldn't set float column %d", i)
		}
	}
	for i, row := range m.Rows {
		if row == nil {
			return errors.Errorf("invalid gpio for row %d", i)
		}
		if err := row.In(gpio.PullUp, gpio.NoEdge); err != nil {
			return errors.Wrapf(err, "couldn't pull up row %d", i)
		}
	}
	return nil
}

func (m *ReedMatrix) Read() ([][]bool, error) {
	data := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		data[i] = make([]bool, 8)
		if err := m.Columns[i].Out(gpio.Low); err != nil {
			return nil, errors.Wrapf(err, "couldn't set low column %d", i)
		}
		time.Sleep(time.Millisecond)
		for j := 0; j < 8; j++ {
			data[i][j] = !bool(m.Rows[j].Read())
		}
		if err := m.Columns[i].In(gpio.Float, gpio.NoEdge); err != nil {
			return nil, errors.Wrapf(err, "couldn't set float column %d", i)
		}
	}
	return data, nil
}
