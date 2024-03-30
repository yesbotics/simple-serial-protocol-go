package ssp

type CommandParam struct {
	paramType ParamType
	value     any
}

type WriteCommandConfig struct {
	commandId     byte
	commandParams []CommandParam
}

func NewWriteCommandConfig(commandId byte) *WriteCommandConfig {
	config := &WriteCommandConfig{
		commandId:     commandId,
		commandParams: []CommandParam{},
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
