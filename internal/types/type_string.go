package types

import (
	"yesbotics/ssp/internal/parser"
)

type TypeString struct {
	BaseType
	isFull bool
}

func (t *TypeString) AddByte(bite byte) {
	if bite == parser.CharNull {
		t.isFull = true
		return
	}
	t.data = append(t.data, bite)
}

func (t *TypeString) GetLength() uint32 {
	return 1
}

func (t *TypeString) GetData() (string, error) {
	return string(t.data), nil
}

func (t *TypeString) GetBuffer(data string) []byte {
	b := []byte(data)
	b = append(b, parser.CharNull)
	return b
}

func (t *TypeString) Reset() {
	t.data = nil
	t.isFull = false
}
