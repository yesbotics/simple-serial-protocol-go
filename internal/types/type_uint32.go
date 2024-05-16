package types

import (
	"encoding/binary"
	"errors"
)

type typeUint32 struct {
	baseType
}

func NewTypeUint32() Type {
	return &typeUint32{
		baseType{
			index:  0,
			length: 4,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeUint32) GetLength() uint32 {
	return 4
}

func (t *typeUint32) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return binary.LittleEndian.Uint32(t.data), nil
}

func (t *typeUint32) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(uint32); ok {
		b := make([]byte, t.GetLength())
		binary.LittleEndian.PutUint32(b, value)
		return b, nil
	} else {
		return nil, errors.New("type assertion to uint32 failed")
	}
}
