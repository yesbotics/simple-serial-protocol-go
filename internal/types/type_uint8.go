package types

import (
	"errors"
)

type typeUint8 struct {
	typeBool
}

func NewTypeUint8() Type {
	return &typeUint8{
		typeBool{
			baseType{
				index:  0,
				length: 1,
				data:   make([]byte, 0),
			},
		},
	}
}

func (t *typeUint8) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return t.data[0], nil
}

func (t *typeUint8) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(uint8); ok {
		return []byte{value}, nil
	} else {
		return nil, errors.New("type assertion to uint8 failed")
	}
}
