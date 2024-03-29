package types

import "errors"

type TypeInt8 struct {
	BaseType
}

func (t *TypeInt8) GetLength() uint32 {
	return 1
}

func (t *TypeInt8) GetData() (int8, error) {
	if !t.IsFull() {
		return 0, errors.New("no data available")
	}
	return int8(t.data[0]), nil
}

func (t *TypeInt8) GetBuffer(data int8) []byte {
	return []byte{byte(data)}
}
