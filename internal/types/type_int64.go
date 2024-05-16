package types

import (
	"encoding/binary"
	"errors"
)

type typeInt64 struct {
	baseType
}

func NewTypeInt64() Type {
	return &typeInt64{
		baseType: baseType{
			index:  0,
			length: 8,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeInt64) GetLength() uint32 {
	return 8
}

func (t *typeInt64) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int64(binary.LittleEndian.Uint64(t.data)), nil
}

func (t *typeInt64) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(int64); ok {
		b := make([]byte, t.GetLength())
		binary.LittleEndian.PutUint64(b, uint64(value))
		return b, nil
	} else {
		return nil, errors.New("type assertion to int64 failed")
	}
}
