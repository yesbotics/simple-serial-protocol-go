package types

import (
	"encoding/binary"
	"errors"
)

type typeInt32 struct {
	baseType
}

func NewTypeInt32() Type {
	return &typeInt32{
		baseType: baseType{
			index:  0,
			length: 4,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeInt32) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int32(binary.LittleEndian.Uint32(t.data)), nil
}

func (t *typeInt32) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(int32); ok {
		b := make([]byte, t.length)
		binary.LittleEndian.PutUint32(b, uint32(value))
		return b, nil
	} else {
		return nil, errors.New("type assertion to int32 failed")
	}
}
