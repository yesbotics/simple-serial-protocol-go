package parser

import "yesbotics/ssp"

type Command struct {
	commandId    byte
	callback     ssp.CommandCallback
	paramsParser *ParamsParser
}

func NewCommand(commandId byte, callback ssp.CommandCallback, paramTypes []ssp.ParamType) *Command {
	return &Command{
		commandId:    commandId,
		callback:     callback,
		paramsParser: NewParamsParser(paramTypes),
	}
}

func (c *Command) ParamsRead() bool {
	if c.paramsParser != nil {
		return c.paramsParser.IsFull()
	}
	return true
}

func (c *Command) AddByte(bite byte) {
	if c.paramsParser != nil {
		_ = c.paramsParser.AddByte(bite)
	}
}

func (c *Command) ResetParamParser() {
	if c.paramsParser != nil {
		c.paramsParser.Reset()
	}
}

func (c *Command) Dispose() {
	c.callback = nil
	if c.paramsParser != nil {
		c.paramsParser.Dispose()
	}
}

func (c *Command) CallCallback() {
	if c.paramsParser != nil {
		c.callback(c.paramsParser.GetData())
	} else {
		c.callback()
	}
}
