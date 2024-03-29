package ssp

type ReadCommandConfig struct {
	commandId byte
	callback  CallbackFunc
	types     []ParamType
}

func NewReadCommandConfig(commandId byte, callback CallbackFunc) *ReadCommandConfig {
	config := &ReadCommandConfig{
		commandId: commandId,
		callback:  callback,
	}

	return config
}

func (r *ReadCommandConfig) GetCallback() CallbackFunc {
	return r.callback
}

func (r *ReadCommandConfig) GetCommandId() byte {
	return r.commandId
}

func (r *ReadCommandConfig) GetCommandParamTypes() []ParamType {
	return r.types
}

func (r *ReadCommandConfig) AddParam(paramType ParamType) *ReadCommandConfig {
	r.types = append(r.types, paramType)
	return r
}
