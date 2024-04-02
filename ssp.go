package ssp

import (
	"encoding/hex"
	"errors"
	"fmt"
	"go.bug.st/serial"
	"log/slog"
	"yesbotics/ssp/internal/parser"
)

type Ssp struct {
	portname       string
	baudrate       int
	commands       map[byte]*parser.Command
	port           *serial.Port
	mode           *serial.Mode
	connected      bool
	currentCommand *parser.Command
}

type CommandCallback func(args ...any)

func NewSsp(portname string, baudrate int) *Ssp {

	mode := &serial.Mode{
		BaudRate: baudrate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	ssp := &Ssp{
		portname:  portname,
		baudrate:  baudrate,
		mode:      mode,
		port:      nil,
		connected: false,
	}

	return ssp
}

func (s *Ssp) Open() error {
	port, err := serial.Open(s.portname, s.mode)
	if err != nil {
		return err
	}

	s.port = &port
	s.connected = true

	go s.readSerialData()

	return nil
}

func (s *Ssp) Close() {
	s.connected = false
	if s.port != nil {
		_ = (*s.port).Close()
	}
	s.port = nil
}

func (s *Ssp) Dispose() {
	s.Close()
	s.commands = make(map[byte]*parser.Command)
}

func (s *Ssp) RegisterCommand(command *ReadCommandConfig) {
	s.commands[command.GetCommandId()] = parser.NewCommand(
		command.GetCommandId(),
		command.GetCallback(),
		command.GetCommandParamTypes(),
	)
}

//func (s *Ssp) RegisterCommand(commandId byte, callback CommandCallback) {
//	s.commands[commandId] = parser.NewCommand(commandId, callback)
//}

func (s *Ssp) UnregisterCommand(commandId byte) {
	command := s.commands[commandId]
	if command != nil {
		command.Dispose()
	}
	delete(s.commands, commandId)
}

func (s *Ssp) WriteCommand(config *WriteCommandConfig) error {
	_, err := s.write([]byte{config.GetCommandId()})
	if err != nil {
		return err
	}

	if config.HasParameters() {
		for _, commandParam := range config.GetCommandParams() {
			_, err = s.write(parser.BufferCeator.GetBuffer(commandParam.paramType, commandParam.value))
			if err != nil {
				return err
			}
		}
	}

	_, err = s.write([]byte{parser.CharEot})
	if err != nil {
		return err
	}

	return nil
}

func (s *Ssp) write(buffer []byte) (int, error) {
	fmt.Printf("W: %s", buffer)
	return (*s.port).Write(buffer)
}

func (s *Ssp) readSerialData() {
	buffer := make([]byte, 128)
	for {
		if !s.connected {
			return
		}

		n, err := (*s.port).Read(buffer)
		if err != nil {
			slog.Error(fmt.Sprintf("\"Could not read serial data: %s", err))
			return
		}

		_ = s.onData(buffer[:n])
	}
}

func (s *Ssp) onData(bytes []byte) error {

	bite := bytes[0]

	if s.currentCommand != nil {
		/**
		 * Got command already -> reading param data
		 */
		if s.currentCommand.ParamsRead() {
			if bite == parser.CharEot {
				s.currentCommand.CallCallback()
				s.currentCommand = nil
			} else {
				return errors.New("EOT byte was not read")
			}
		} else {
			s.currentCommand.AddByte(bite)
		}
	} else {
		command, ok := s.commands[bite]
		if !ok {
			return errors.New("Command not found: " + hex.EncodeToString([]byte{bite}))
		} else {
			s.currentCommand = command
			s.currentCommand.ResetParamParser()
		}
	}
	return nil
}
