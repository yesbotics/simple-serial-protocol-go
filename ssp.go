package ssp

import (
	"fmt"
	"yesbotics/ssp/internal/parser"
)

type Ssp struct {
	portname string
	baudrate int32
	commands map[byte]*Command
	//paramTypeInstances map[ParamType]*types.Type
}

type Command struct {
	commandId byte
	callback  CallbackFunc
}

type CallbackFunc func(args ...interface{})

const (
	CharEot  byte = 0x0A // End of Transmission - Line Feed Zeichen \n
	CharNull byte = 0x00 // End of String
)

func New(portname string, baudrate int32) *Ssp {

	ssp := &Ssp{
		portname: portname,
		baudrate: baudrate,
	}

	ssp.initParamTypes()

	return ssp
}

func (s *Ssp) RegisterCommand(commandId byte, callback CallbackFunc) {
	s.commands[commandId] = &Command{commandId: commandId, callback: callback}
}

func (s *Ssp) UnregisterCommand(commandId byte) {
	delete(s.commands, commandId)
}

func (s *Ssp) WriteCommand(config WriteCommandConfig) {
	s.write([]byte{config.GetCommandId()})

	if config.HasParameters() {
		for _, commandParam := range config.GetCommandParams() {
			parser.BufferCeator.GetBuffer(commandParam.paramType, commandParam.value)
		}
	}

	s.write([]byte{CharEot})
}

func (s *Ssp) initParamTypes() {

}

func (s *Ssp) addParamType(name ParamType) {

}

func (s *Ssp) write(buffer []byte) {
	fmt.Printf("W: %s", buffer)
	//this.serialPort.write(buffer, "ascii");
}
