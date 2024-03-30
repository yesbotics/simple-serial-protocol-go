package types

import (
	"encoding/binary"
	"errors"
)

type TypeUint16 struct {
	BaseType
}

func (t *TypeUint16) GetLength() uint32 {
	return 2
}

func (t *TypeUint16) GetData() (uint16, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return binary.LittleEndian.Uint16(t.data), nil
}

func (t *TypeUint16) GetBuffer(data uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, data)
	return b
}
