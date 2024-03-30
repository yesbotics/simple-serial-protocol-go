package types

import (
	"encoding/binary"
	"errors"
)

type TypeInt64 struct {
	BaseType
}

func (t *TypeInt64) GetLength() uint32 {
	return 8
}

func (t *TypeInt64) GetData() (int64, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int64(binary.LittleEndian.Uint64(t.data)), nil
}

func (t *TypeInt64) GetBuffer(data int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(data))
	return b
}
