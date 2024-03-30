package types

import "errors"

type TypeBool struct {
	BaseType
}

func (t *TypeBool) GetLength() uint32 {
	return 1
}

func (t *TypeBool) GetData() (bool, error) {
	if !t.IsFull() {
		return false, errors.New("no data available")
	}
	return t.data[0] == 1, nil
}

func (t *TypeBool) GetBuffer(data bool) []byte {
	if data {
		return []byte{1}
	} else {
		return []byte{0}
	}
}
