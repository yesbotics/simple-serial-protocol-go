package types

//type GenType[T byte | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | bool | float32 | string] struct {
//	data   []byte
//	index  uint32
//	length uint32
//}
//
//func New[T byte | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | bool | float32 | string]() *GenType[T] {
//
//	//var test T
//	var length uint32 = 0
//
//	//switch hmm := test.(type) {
//	//
//	//}
//
//	return &GenType[T]{
//		data:   []byte{},
//		index:  0,
//		length: length,
//	}
//}
//
//func (t *GenType[T]) Reset() {
//	t.index = 0
//}
//
//func (t *GenType[T]) AddByte(bite byte) {
//	t.data = append(t.data, bite)
//}
//
//func (t *GenType[T]) IsFull() bool {
//	return t.index >= t.GetLength()
//}
//
//func (t *GenType[T]) GetLength() uint32 {
//	return t.length
//}
//
//func (t *GenType[T]) Dispose() {
//
//}
//
//func (t *GenType[T]) GetData() (T, error) {
//	if !t.IsFull() {
//		return nil, errors.New("no data available")
//	}
//	return t.data[0], nil
//}
//
//func (t *GenType[T]) GetBuffer(data uint8) []byte {
//	return []byte{data}
//}
