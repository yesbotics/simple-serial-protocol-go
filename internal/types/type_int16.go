package types

import (
	"encoding/binary"
	"errors"
)

type typeInt16 struct {
	baseType
}

func NewTypeInt16() Type {
	return &typeInt16{
		baseType: baseType{
			index:  0,
			length: 2,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeInt16) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int16(binary.LittleEndian.Uint16(t.data)), nil
}

func (t *typeInt16) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(int16); ok {
		b := make([]byte, t.length)
		binary.LittleEndian.PutUint16(b, uint16(value))
		return b, nil
	} else {
		return nil, errors.New("type assertion to int16 failed")
	}
}
