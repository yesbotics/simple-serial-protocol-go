package types

import (
	"encoding/binary"
	"errors"
)

type TypeUint64 struct {
	BaseType
}

func (t *TypeUint64) GetLength() uint32 {
	return 8
}

func (t *TypeUint64) GetData() (uint64, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return binary.LittleEndian.Uint64(t.data), nil
}

func (t *TypeUint64) GetBuffer(data uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, data)
	return b
}
