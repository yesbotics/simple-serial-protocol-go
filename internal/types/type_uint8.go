package types

import "errors"

type TypeUInt8 struct {
	TypeInt8
}

func (t *TypeUInt8) GetData() (uint8, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return t.data[0], nil
}

func (t *TypeUInt8) GetBuffer(data uint8) []byte {
	return []byte{data}
}
