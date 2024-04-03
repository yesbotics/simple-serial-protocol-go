package parser

import (
	"yesbotics/ssp/internal/types"
	"yesbotics/ssp/pkg/config"
)

type bufferCreator struct {
	types map[config.ParamType]any
}

var BufferCeator bufferCreator = newBufferCreator()

func newBufferCreator() bufferCreator {

	creator := bufferCreator{
		types: make(map[config.ParamType]any),
	}

	creator.types[config.ParamTypeBool] = types.TypeBool{}
	creator.types[config.ParamTypeByte] = types.TypeByte{}
	creator.types[config.ParamTypeInt8] = types.TypeInt8{}
	creator.types[config.ParamTypeInt16] = types.TypeInt16{}
	creator.types[config.ParamTypeInt32] = types.TypeInt32{}
	creator.types[config.ParamTypeInt64] = types.TypeInt64{}
	creator.types[config.ParamTypeUint8] = types.TypeUint8{}
	creator.types[config.ParamTypeUint16] = types.TypeUint16{}
	creator.types[config.ParamTypeUint32] = types.TypeUint32{}
	creator.types[config.ParamTypeUint64] = types.TypeUint64{}
	creator.types[config.ParamTypeFloat32] = types.TypeFloat32{}
	creator.types[config.ParamTypeFloat64] = types.TypeFloat64{}
	creator.types[config.ParamTypeString] = types.TypeString{}

	return creator
}

func (b *bufferCreator) GetBuffer(paramType config.ParamType, data any) []byte {
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
