package types

import "errors"

type typeString struct {
	baseType
	isFull bool
}

func NewTypeString() Type {
	return &typeString{
		baseType: baseType{
			index:  0,
			length: 0,
			data:   make([]byte, 0),
		},
		isFull: false,
	}
}

func (t *typeString) AddByte(bite byte) {
	if t.isFull {
		return
	}
	if bite == CharNull {
		t.isFull = true
		return
	}
	t.data = append(t.data, bite)
}

func (t *typeString) IsFull() bool {
	return t.isFull
}

func (t *typeString) GetData() (any, error) {
	return string(t.data), nil
}

func (t *typeString) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(string); ok {
		b := []byte(value)
		b = append(b, CharNull)
		return b, nil
	} else {
		return nil, errors.New("type assertion to string failed")
	}
}

func (t *typeString) Reset() {
	t.data = make([]byte, 0)
	t.isFull = false
}
