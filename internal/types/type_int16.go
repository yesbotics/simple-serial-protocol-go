package types

import (
	"encoding/binary"
	"errors"
)

type TypeInt16 struct {
	BaseType
}

func (t *TypeInt16) GetLength() uint32 {
	return 2
}

func (t *TypeInt16) GetData() (int16, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int16(binary.LittleEndian.Uint16(t.data)), nil
}

func (t *TypeInt16) GetBuffer(data int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(data))
	return b
}
