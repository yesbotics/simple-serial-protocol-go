package simple_serial_protocol

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/yesbotics/simple-serial-protocol-go/internal/parser"
	"github.com/yesbotics/simple-serial-protocol-go/internal/types"
	"github.com/yesbotics/simple-serial-protocol-go/pkg/config"
	"go.bug.st/serial"
	"log/slog"
)

type SimpleSerialProtocol struct {
	portname       string
	baudrate       int
	commands       map[byte]*parser.Command
	port           *serial.Port
	mode           *serial.Mode
	connected      bool
	currentCommand *parser.Command
}

func NewSimpleSerialProtocol(portname string, baudrate int) *SimpleSerialProtocol {

	mode := &serial.Mode{
		BaudRate: baudrate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	return &SimpleSerialProtocol{
		portname:       portname,
		baudrate:       baudrate,
		commands:       make(map[byte]*parser.Command),
		port:           nil,
		mode:           mode,
		connected:      false,
		currentCommand: nil,
	}
}

func (s *SimpleSerialProtocol) Open() error {
	if s.connected {
		return errors.New("already connected")
	}

	port, err := serial.Open(s.portname, s.mode)
	if err != nil {
		return err
	}

	s.port = &port
	s.connected = true

	go s.readSerialData()

	return nil
}

func (s *SimpleSerialProtocol) Close() {
	s.connected = false
	if s.port != nil {
		_ = (*s.port).Close()
	}
	s.port = nil
}

func (s *SimpleSerialProtocol) Dispose() {
	s.Close()
	s.commands = make(map[byte]*parser.Command)
}

func (s *SimpleSerialProtocol) RegisterCommand(command *config.ReadCommandConfig) {
	s.commands[command.GetCommandId()] = parser.NewCommand(
		command.GetCommandId(),
		command.GetCallback(),
		command.GetCommandParamTypes(),
	)
}

//func (s *SimpleSerialProtocol) RegisterCommand(commandId byte, callback CommandCallback) {
//	s.commands[commandId] = parser.NewCommand(commandId, callback)
//}

func (s *SimpleSerialProtocol) UnregisterCommand(commandId byte) {
	command := s.commands[commandId]
	if command != nil {
		command.Dispose()
	}
	delete(s.commands, commandId)
}

func (s *SimpleSerialProtocol) WriteCommand(config *config.WriteCommandConfig) error {
	_, err := s.write([]byte{config.GetCommandId()})
	if err != nil {
		return err
	}

	if config.HasParameters() {
		for _, commandParam := range config.GetCommandParams() {
			_, err = s.write(parser.BufferCeator.GetBuffer(commandParam.ParamType, commandParam.Value))
			if err != nil {
				return err
			}
		}
	}

	_, err = s.write([]byte{types.CharEot})
	if err != nil {
		return err
	}

	return nil
}

func (s *SimpleSerialProtocol) write(buffer []byte) (int, error) {
	slog.Debug("Write", "data", fmt.Sprintf("%#x", buffer))
	return (*s.port).Write(buffer)
}

func (s *SimpleSerialProtocol) readSerialData() {
	buffer := make([]byte, 10)
	for {
		if !s.connected {
			return
		}

		n, err := (*s.port).Read(buffer)
		if err != nil {
			slog.Error("\"Could not read serial data", slog.Any("data", err))
			return
		}

		_ = s.onData(buffer[:n])
	}

}

func (s *SimpleSerialProtocol) onData(bytes []byte) error {
	for _, bite := range bytes {
		if s.currentCommand != nil {
			/**
			 * Got command already -> reading param data
			 */
			if s.currentCommand.ParamsRead() {
				if bite == types.CharEot {
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
	}
	return nil
}
