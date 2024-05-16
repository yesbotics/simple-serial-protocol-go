package types

const (
	CharEot  byte = 0x0A // End of Transmission - Line Feed Zeichen \n
	CharNull byte = 0x00 // End of String
)

type Type interface {
	Reset()
	AddByte(bite byte)
	IsFull() bool
	GetData() (any, error)
	Dispose()
	GetBuffer(data any) ([]byte, error)
	//GetLength() uint32
}

//type Type[T int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | bool | float32 | string] interface {
//	Reset()
//	AddByte(bite byte)
//	IsFull() bool
//	GetData() (T, error)
//	Dispose()
//	GetBuffer(data T) []byte
//	GetLength() uint32
//}

//func GetBuffer[T byte | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | bool | float32 | string](paramType ssp.ParamType, data T) []byte {
//	b := make([]byte, 2)
//	binary.LittleEndian.PutUint16(b, data)
//	return b
//}
