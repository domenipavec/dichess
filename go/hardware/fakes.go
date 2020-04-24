package hardware

import (
	"bufio"
	"bytes"
	"io/ioutil"
)

type FakeCoil struct{}

func (c *FakeCoil) Initialize() error {
	return nil
}

func (c *FakeCoil) On() error {
	return nil
}

func (c *FakeCoil) Off() error {
	return nil
}

func (c *FakeCoil) Rotate(v int) error {
	return nil
}

func (c *FakeCoil) SetPwm(v uint8) error {
	return nil
}

type FakeAxis struct{}

func (a *FakeAxis) Initialize() error {
	return nil
}

func (a *FakeAxis) SetCurrent(v uint8) error {
	return nil
}

func (a *FakeAxis) GoTo(t float64, s uint8) error {
	return nil
}

type FakeMatrix struct{}

func (m *FakeMatrix) Initialize() error {
	data := `01111111
11111111
00000000
00000000
00000000
00000000
11111111
11111111`
	return ioutil.WriteFile("chessboard", []byte(data), 0o666)
}

func (m *FakeMatrix) Read() ([][]bool, error) {
	data, err := ioutil.ReadFile("chessboard")
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	result := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		for j := 0; j < 8; j++ {
			if i == 0 {
				result[j] = make([]bool, 8)
			}
			result[j][i] = line[j] == '1'
		}
	}
	return result, nil
}
