package types

import (
	"errors"
)

type typeInt8 struct {
	baseType
}

func NewTypeInt8() Type {
	return &typeInt8{
		baseType: baseType{
			index:  0,
			length: 1,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeInt8) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int8(t.data[0]), nil
}

func (t *typeInt8) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(int8); ok {
		return []byte{byte(value)}, nil
	} else {
		return nil, errors.New("type assertion to float32 failed")
	}
}
