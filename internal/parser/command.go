package parser

import (
	"yesbotics/ssp/pkg/config"
)

type Command struct {
	commandId    byte
	callback     config.CommandCallback
	paramsParser *ParamsParser
}

func NewCommand(commandId byte, callback config.CommandCallback, paramTypes []config.ParamType) *Command {
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
