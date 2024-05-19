package parser

import (
	"github.com/yesbotics/simple-serial-protocol-go/internal/types"
	"github.com/yesbotics/simple-serial-protocol-go/pkg/config"
)

type bufferCreator struct {
	types map[config.ParamType]types.Type
}

var BufferCeator = newBufferCreator()

func newBufferCreator() bufferCreator {

	creator := bufferCreator{
		types: make(map[config.ParamType]types.Type),
	}

	creator.types[config.ParamTypeBool] = types.NewTypeBool()
	creator.types[config.ParamTypeByte] = types.NewTypeByte()
	creator.types[config.ParamTypeChar] = types.NewTypeChar()
	creator.types[config.ParamTypeInt8] = types.NewTypeInt8()
	creator.types[config.ParamTypeInt16] = types.NewTypeInt16()
	creator.types[config.ParamTypeInt32] = types.NewTypeInt32()
	creator.types[config.ParamTypeInt64] = types.NewTypeInt64()
	creator.types[config.ParamTypeUint8] = types.NewTypeUint8()
	creator.types[config.ParamTypeUint16] = types.NewTypeUint16()
	creator.types[config.ParamTypeUint32] = types.NewTypeUint32()
	creator.types[config.ParamTypeUint64] = types.NewTypeUint64()
	creator.types[config.ParamTypeFloat32] = types.NewTypeFloat32()
	creator.types[config.ParamTypeFloat64] = types.NewTypeFloat64()
	creator.types[config.ParamTypeString] = types.NewTypeString()

	return creator
}

func (b *bufferCreator) GetBuffer(paramType config.ParamType, data any) []byte {
	theType := b.types[paramType]

	buffer, err := theType.GetBuffer(data)
	if err != nil {
		return []byte{}
	}

	return buffer
}
