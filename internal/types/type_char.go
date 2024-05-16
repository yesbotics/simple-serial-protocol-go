package types

type typeChar struct {
	typeByte
}

func NewTypeChar() Type {
	return &typeChar{
		typeByte: typeByte{
			typeUint8{
				typeBool: typeBool{
					baseType: baseType{
						index:  0,
						length: 1,
						data:   make([]byte, 0),
					},
				},
			},
		},
	}
}
