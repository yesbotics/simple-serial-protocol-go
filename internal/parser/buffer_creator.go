package parser

import (
	"yesbotics/ssp"
	"yesbotics/ssp/internal/types"
)

type bufferCreator struct {
	types map[ssp.ParamType]interface{}
}

var BufferCeator bufferCreator = newBufferCreator()

func newBufferCreator() bufferCreator {

	creator := bufferCreator{
		//types: map[ssp.ParamType]*interface{}{},
	}

	creator.types[ssp.ParamTypeByte] = types.TypeByte{}
	creator.types[ssp.ParamTypeInt8] = types.TypeByte{}
	creator.types[ssp.ParamTypeInt16] = types.TypeByte{}
	creator.types[ssp.ParamTypeUint8] = types.TypeByte{}
	creator.types[ssp.ParamTypeUint16] = types.TypeByte{}

	return creator
}

// func (b *bufferCreator) getBuffer[T byte | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | bool | float32 | string](paramType ssp.ParamType, data T) []byte {
func (b *bufferCreator) GetBuffer(paramType ssp.ParamType, data any) []byte {
	theType := b.types[paramType]

	switch inst := theType.(type) {
	case types.TypeByte:
		return inst.GetBuffer(data.(uint8))
	case types.TypeChar:
		return inst.GetBuffer(data.(byte))
	case types.TypeBool:
		return inst.GetBuffer(data.(bool))
	case types.TypeUint8:
		return inst.GetBuffer(data.(uint8))
	case types.TypeInt16:
		return inst.GetBuffer(data.(int16))
	case types.TypeUint16:
		return inst.GetBuffer(data.(uint16))
	case types.TypeInt32:
		return inst.GetBuffer(data.(int32))
	case types.TypeUint32:
		return inst.GetBuffer(data.(uint32))
	case types.TypeInt64:
		return inst.GetBuffer(data.(int64))
	case types.TypeUint64:
		return inst.GetBuffer(data.(uint64))
	case types.TypeFloat32:
		return inst.GetBuffer(data.(float32))
	case types.TypeFloat64:
		return inst.GetBuffer(data.(float64))
	case types.TypeString:
		return inst.GetBuffer(data.(string))
	}

	return []byte{}
}
