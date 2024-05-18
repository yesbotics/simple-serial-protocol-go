package types

import (
	"encoding/binary"
	"errors"
	"math"
)

type typeFloat32 struct {
	baseType
}

func NewTypeFloat32() Type {
	return &typeFloat32{
		baseType: baseType{
			index:  0,
			length: 4,
			data:   make([]byte, 0),
		},
	}
}

func (t *typeFloat32) GetData() (any, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(t.data)), nil
}

func (t *typeFloat32) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(float32); ok {
		b := make([]byte, t.length)
		binary.LittleEndian.PutUint32(b, math.Float32bits(value))
		return b, nil
	} else {
		return nil, errors.New("type assertion to float32 failed")
	}
}
