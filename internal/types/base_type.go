package types

type BaseType struct {
	data   []byte
	index  uint32
	length uint32
}

func (b *BaseType) Reset() {
	b.index = 0
}

func (b *BaseType) AddByte(bite byte) {
	b.data = append(b.data, bite)
}

func (b *BaseType) IsFull() bool {
	return b.index >= b.GetLength()
}

func (b *BaseType) GetLength() uint32 {
	return 1
}

func (b *BaseType) Dispose() {

}
