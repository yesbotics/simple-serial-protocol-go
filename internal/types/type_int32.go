package types

import (
	"encoding/binary"
	"errors"
)

type TypeInt32 struct {
	BaseType
}

func (t *TypeInt32) GetLength() uint32 {
	return 4
}

func (t *TypeInt32) GetData() (int32, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int32(binary.LittleEndian.Uint32(t.data)), nil
}

func (t *TypeInt32) GetBuffer(data int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(data))
	return b
}
