package types

//type Type interface {
//	Reset()
//	AddByte(bite byte)
//	IsFull() bool
//	GetData() (any, error)
//	Dispose()
//	GetBuffer(data any) []byte
//	GetLength() uint32
//}

type Type[T byte | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | bool | float32 | string] interface {
	Reset()
	AddByte(bite byte)
	IsFull() bool
	GetData() (T, error)
	Dispose()
	GetBuffer(data T) []byte
	GetLength() uint32
}
