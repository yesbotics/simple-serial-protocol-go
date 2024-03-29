package ssp

import "fmt"

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
	CharEot byte = 0x0A // End of Transmission - Line Feed Zeichen \n
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
	if !config.HasParameters() {
		return
	}

	for _, commandParam := range config.GetCommandParams() {

	}
}

func (s *Ssp) initParamTypes() {

}

func (s *Ssp) addParamType(name ParamType) {

}

func (s *Ssp) write(buffer []byte) {
	fmt.Printf("W: %s", buffer)
	//this.serialPort.write(buffer, "ascii");
}
