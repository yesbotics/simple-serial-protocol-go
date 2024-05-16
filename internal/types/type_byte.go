package types

import "errors"

type typeByte struct {
	typeUint8
}

func NewTypeByte() Type {
	return &typeByte{
		typeUint8: typeUint8{
			typeBool: typeBool{
				baseType: baseType{
					index:  0,
					length: 1,
					data:   make([]byte, 0),
				},
			},
		},
	}
}

func (t *typeByte) GetBuffer(data any) ([]byte, error) {
	if value, ok := data.(byte); ok {
		return []byte{value}, nil
	} else {
		return nil, errors.New("type assertion to uint8 failed")
	}
}
