package types

import "errors"

type TypeUint8 struct {
	TypeBool
}

func (t *TypeUint8) GetData() (uint8, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return t.data[0], nil
}

func (t *TypeUint8) GetBuffer(data uint8) []byte {
	return []byte{data}
}
