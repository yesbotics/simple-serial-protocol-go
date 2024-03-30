package types

import (
	"encoding/binary"
	"errors"
)

type TypeFloat32 struct {
	BaseType
}

func (t *TypeFloat32) GetLength() uint32 {
	return 4
}

func (t *TypeFloat32) GetData() (float32, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return float32(binary.LittleEndian.Uint32(t.data)), nil
}

func (t *TypeFloat32) GetBuffer(data float32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(data))
	return b
}
