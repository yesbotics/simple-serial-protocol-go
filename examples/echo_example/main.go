package main

import (
	"fmt"
	"yesbotics/ssp/pkg/config"
	"yesbotics/ssp/pkg/simple_serial_protocol"
)

func main() {

	fmt.Println("Starting echo example.")

	portname := "/dev/ttyUSB0"
	baudrate := 57600

	arduino := simple_serial_protocol.NewSsp(portname, baudrate)
	readConfig := config.NewReadCommandConfig(byte('s'), onRead)
	readConfig.AddParam(config.ParamTypeByte)
	readConfig.AddParam(config.ParamTypeBool)
	readConfig.AddParam(config.ParamTypeInt8)
	readConfig.AddParam(config.ParamTypeUint8)
	readConfig.AddParam(config.ParamTypeInt16)
	readConfig.AddParam(config.ParamTypeUint16)
	readConfig.AddParam(config.ParamTypeInt32)
	readConfig.AddParam(config.ParamTypeUint32)
	readConfig.AddParam(config.ParamTypeInt64)
	readConfig.AddParam(config.ParamTypeUint64)
	readConfig.AddParam(config.ParamTypeFloat32)
	readConfig.AddParam(config.ParamTypeChar)
	readConfig.AddParam(config.ParamTypeString)
	readConfig.AddParam(config.ParamTypeString)
	readConfig.AddParam(config.ParamTypeString)

	arduino.RegisterCommand(readConfig)
	err := arduino.Open()
	if err != nil {
		fmt.Printf("Could not connect. Error %q\n", err)
		return
	}

	write := config.NewWriteCommandConfig(
		byte('s'),
		[]any{
			byte(0xff),
			true,
			int8(-128),
			uint8(255),
			int16(-32768),
			uint16(65523),
			int32(-2147483648),
			uint32(4294967295),

			// For BigInt Support add `esnext` to your tsconfig lib section
			int64(-2147483648000999),
			uint64(7294967295000999),

			float32(-1.23456789101112),
			byte('J'),

			"text1: Hey, I'm text one!",
			"text2: And I am his brother text two!",
			"text3: Nice!",
		},
	)

	err = arduino.WriteCommand(write)
	if err != nil {
		fmt.Printf("Could not write command: %s\n", err)
		return
	}

	fmt.Println("nice")
}

func onRead(args ...any) {

	fmt.Println("Received several values from Arduino:")
	fmt.Printf("byteValue: %s \n", args[0])
	fmt.Printf("booleanValue: %s \n", args[1])
	fmt.Printf("int8Value: %s \n", args[2])
	fmt.Printf("uint8Value: %s \n", args[3])
	fmt.Printf("int16Value: %s \n", args[4])
	fmt.Printf("uint16Value: %s \n", args[5])
	fmt.Printf("int32Value: %s \n", args[6])
	fmt.Printf("uint32Value: %s \n", args[7])
	fmt.Printf("int64Value: %s \n", args[8])
	fmt.Printf("uint64Value: %s \n", args[9])
	fmt.Printf("floatValue: %s \n", args[10])
	fmt.Printf("charValue: %s \n", args[11])
	fmt.Printf("stringValue1: %s \n", args[12])
	fmt.Printf("stringValue2: %s \n", args[13])
	fmt.Printf("stringValue3: %s \n", args[14])

	fmt.Printf("Read data: %s", args)
}
