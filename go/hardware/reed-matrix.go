package hardware

import (
	"github.com/pkg/errors"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

type ReedMatrix struct {
	Columns []gpio.PinIO
	Rows    []gpio.PinIn
}

func NewMatrix() (*ReedMatrix, error) {
	m := &ReedMatrix{
		Columns: []gpio.PinIO{
			gpioreg.ByName("GPIO4"),
			gpioreg.ByName("GPIO17"),
			gpioreg.ByName("GPIO27"),
			gpioreg.ByName("GPIO22"),
			gpioreg.ByName("GPIO5"),
			gpioreg.ByName("GPIO6"),
			gpioreg.ByName("GPIO13"),
			gpioreg.ByName("GPIO19"),
		},
		Rows: []gpio.PinIn{
			gpioreg.ByName("GPIO18"),
			gpioreg.ByName("GPIO23"),
			gpioreg.ByName("GPIO24"),
			gpioreg.ByName("GPIO25"),
			gpioreg.ByName("GPIO12"),
			gpioreg.ByName("GPIO16"),
			gpioreg.ByName("GPIO20"),
			gpioreg.ByName("GPIO21"),
		},
	}
	for i, col := range m.Columns {
		if col == nil {
			return nil, errors.Errorf("invalid gpio for column %v", i)
		}
		if err := col.In(gpio.Float, gpio.NoEdge); err != nil {
			return nil, errors.Wrapf(err, "couldn't set float column %v", i)
		}
	}
	for i, row := range m.Rows {
		if row == nil {
			return nil, errors.Errorf("invalid gpio for row %v", i)
		}
		if err := row.In(gpio.PullUp, gpio.NoEdge); err != nil {
			return nil, errors.Wrapf(err, "couldn't pull up row %v", i)
		}
	}
	return m, nil
}

func (m *ReedMatrix) Read() ([][]bool, error) {
	data := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		data[i] = make([]bool, 8)
		if err := m.Columns[i].Out(gpio.Low); err != nil {
			return nil, errors.Wrapf(err, "couldn't set low column %v", i)
		}
		for j := 0; j < 8; j++ {
			data[i][j] = !bool(m.Rows[j].Read())
		}
		if err := m.Columns[i].In(gpio.Float, gpio.NoEdge); err != nil {
			return nil, errors.Wrapf(err, "couldn't set float column %v", i)
		}
	}
	return data, nil
}
