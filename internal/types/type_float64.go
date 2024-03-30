package types

import (
	"encoding/binary"
	"errors"
)

type TypeFloat64 struct {
	BaseType
}

func (t *TypeFloat64) GetLength() uint32 {
	return 8
}

func (t *TypeFloat64) GetData() (float64, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return float64(binary.LittleEndian.Uint64(t.data)), nil
}

func (t *TypeFloat64) GetBuffer(data float64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(data))
	return b
}
