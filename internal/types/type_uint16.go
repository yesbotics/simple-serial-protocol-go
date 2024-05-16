package types

import (
	"encoding/binary"
	"errors"
)

type typeUint16 struct {
	baseType
}

func NewTypeUint16() Type {
	return &typeUint16{
		baseType{
			index:  0,
			length: 2,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeUint16) GetLength() uint32 {
	return 2
}

func (t *typeUint16) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return binary.LittleEndian.Uint16(t.data), nil
}

func (t *typeUint16) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(uint16); ok {
		b := make([]byte, t.GetLength())
		binary.LittleEndian.PutUint16(b, value)
		return b, nil
	} else {
		return nil, errors.New("type assertion to uint16 failed")
	}
}
