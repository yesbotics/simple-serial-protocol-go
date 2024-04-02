package ssp

type CommandParam struct {
	paramType ParamType
	value     any
}

type WriteCommandConfig struct {
	commandId     byte
	commandParams []CommandParam
}

func NewWriteCommandConfig(commandId byte, commandParams []any) *WriteCommandConfig {

	var params []CommandParam

	for _, commandParam := range commandParams {
		switch commandParam.(type) {
		case byte:
			params = append(params, CommandParam{ParamTypeByte, commandParam})
		case bool:
			params = append(params, CommandParam{ParamTypeBool, commandParam})
		case int8:
			params = append(params, CommandParam{ParamTypeInt8, commandParam})
		//case uint8:
		//	params = append(params, CommandParam{ParamTypeUint8, commandParam})
		case int16:
			params = append(params, CommandParam{ParamTypeInt16, commandParam})
		case uint16:
			params = append(params, CommandParam{ParamTypeUint16, commandParam})
		case int32:
			params = append(params, CommandParam{ParamTypeInt32, commandParam})
		case uint32:
			params = append(params, CommandParam{ParamTypeUint32, commandParam})
		case int64:
			params = append(params, CommandParam{ParamTypeInt64, commandParam})
		case uint64:
			params = append(params, CommandParam{ParamTypeUint64, commandParam})
		case float32:
			params = append(params, CommandParam{ParamTypeFloat32, commandParam})
		case float64:
			params = append(params, CommandParam{ParamTypeFloat64, commandParam})
		//case ssp.ParamTypeChar:
		//	params = append(params, &(types.TypeChar{}))
		case string:
			params = append(params, CommandParam{ParamTypeString, commandParam})
		}
	}

	config := &WriteCommandConfig{
		commandId:     commandId,
		commandParams: params,
	}

	return config
}

func (r *WriteCommandConfig) GetCommandId() byte {
	return r.commandId
}

func (r *WriteCommandConfig) GetCommandParams() []CommandParam {
	return r.commandParams
}

func (r *WriteCommandConfig) AddByteValue(value byte) *WriteCommandConfig {
	r.commandParams = append(r.commandParams, CommandParam{
		paramType: ParamTypeByte,
		value:     value,
	})
	return r
}
func (r *WriteCommandConfig) addBooleanValue(value bool) *WriteCommandConfig {
	r.commandParams = append(r.commandParams, CommandParam{
		paramType: ParamTypeBool,
		value:     value,
	})
	return r
}
func (r *WriteCommandConfig) addInt8Value(value int8) *WriteCommandConfig {
	r.commandParams = append(r.commandParams, CommandParam{
		paramType: ParamTypeInt8,
		value:     value,
	})
	return r
}

func (r *WriteCommandConfig) HasParameters() bool {
	return len(r.commandParams) > 0
}
