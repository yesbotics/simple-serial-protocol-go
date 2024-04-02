package ssp

type ReadCommandConfig struct {
	commandId byte
	callback  CommandCallback
	types     []ParamType
}

func NewReadCommandConfig(commandId byte, callback CommandCallback) *ReadCommandConfig {
	config := &ReadCommandConfig{
		commandId: commandId,
		callback:  callback,
	}

	return config
}

func (r *ReadCommandConfig) GetCallback() CommandCallback {
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
