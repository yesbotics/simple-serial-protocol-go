package types

type baseType struct {
	data   []byte
	index  uint32
	length uint32
}

func (b *baseType) Reset() {
	b.index = 0
}

func (b *baseType) AddByte(bite byte) {
	b.data = append(b.data, bite)
	b.index++
}

func (b *baseType) IsFull() bool {
	return b.index >= b.length
}

func (b *baseType) Dispose() {

}
