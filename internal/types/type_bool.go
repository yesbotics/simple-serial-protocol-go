package types

import (
	"errors"
)

type typeBool struct {
	baseType
}

func NewTypeBool() Type {
	return &typeBool{
		baseType: baseType{
			index:  0,
			length: 1,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeBool) GetLength() uint32 {
	return 1
}

func (t *typeBool) GetData() (any, error) {
	if !t.IsFull() {
		return false, errors.New("no data available")
	}
	return t.data[0] == 1, nil
}

func (t *typeBool) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(bool); ok {
		if value {
			return []byte{1}, nil
		} else {
			return []byte{0}, nil
		}
	} else {
		return nil, errors.New("type assertion to bool failed")
	}
}
