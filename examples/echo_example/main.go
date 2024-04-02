package main

import (
	"fmt"
	"yesbotics/ssp"
)

func main() {

	fmt.Println("nice")

	portname := "/dev/ttyUSB0"
	baudrate := 57600

	arduino := ssp.NewSsp(portname, baudrate)
	readConfig := ssp.NewReadCommandConfig(byte('s'), onRead)
	readConfig.AddParam(ssp.ParamTypeByte)
	readConfig.AddParam(ssp.ParamTypeBool)
	readConfig.AddParam(ssp.ParamTypeInt8)
	readConfig.AddParam(ssp.ParamTypeUint8)
	readConfig.AddParam(ssp.ParamTypeInt16)
	readConfig.AddParam(ssp.ParamTypeUint16)
	readConfig.AddParam(ssp.ParamTypeInt32)
	readConfig.AddParam(ssp.ParamTypeUint32)
	readConfig.AddParam(ssp.ParamTypeInt64)
	readConfig.AddParam(ssp.ParamTypeUint64)
	readConfig.AddParam(ssp.ParamTypeFloat32)
	readConfig.AddParam(ssp.ParamTypeChar)
	readConfig.AddParam(ssp.ParamTypeString)
	readConfig.AddParam(ssp.ParamTypeString)
	readConfig.AddParam(ssp.ParamTypeString)

	arduino.RegisterCommand(readConfig)
	err := arduino.Open()
	if err != nil {
		fmt.Printf("Could not connect. Error %q\n", err)
		return
	}

	write := ssp.NewWriteCommandConfig(byte('s'), nil)

	err = arduino.WriteCommand(write)
	if err != nil {
		fmt.Printf("Could not write command: %s\n", err)
		return
	}

	fmt.Println("nice")
}

func onRead(args ...any) {
	fmt.Printf("Read data: %s", args)
}
