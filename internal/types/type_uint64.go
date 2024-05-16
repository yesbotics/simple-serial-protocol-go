package types

import (
	"encoding/binary"
	"errors"
)

type typeUint64 struct {
	baseType
}

func NewTypeUint64() Type {
	return &typeUint64{
		baseType{
			index:  0,
			length: 8,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeUint64) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return binary.LittleEndian.Uint64(t.data), nil
}

func (t *typeUint64) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(uint64); ok {
		b := make([]byte, t.length)
		binary.LittleEndian.PutUint64(b, value)
		return b, nil
	} else {
		return nil, errors.New("type assertion to uint64 failed")
	}
}
