package types

import (
	"encoding/binary"
	"errors"
)

type typeFloat64 struct {
	baseType
}

func NewTypeFloat64() Type {
	return &typeFloat64{
		baseType: baseType{
			index:  0,
			length: 8,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeFloat64) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return float64(binary.LittleEndian.Uint64(t.data)), nil
}

func (t *typeFloat64) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(float64); ok {
		b := make([]byte, t.length)
		binary.LittleEndian.PutUint64(b, uint64(value))
		return b, nil
	} else {
		return nil, errors.New("type assertion to float64 failed")
	}
}
