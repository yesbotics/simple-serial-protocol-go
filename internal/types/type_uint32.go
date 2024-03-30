package types

import (
	"encoding/binary"
	"errors"
)

type TypeUint32 struct {
	BaseType
}

func (t *TypeUint32) GetLength() uint32 {
	return 4
}

func (t *TypeUint32) GetData() (uint32, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return binary.LittleEndian.Uint32(t.data), nil
}

func (t *TypeUint32) GetBuffer(data uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, data)
	return b
}
