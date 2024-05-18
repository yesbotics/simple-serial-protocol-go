package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"yesbotics/ssp/pkg/config"
	"yesbotics/ssp/pkg/simple_serial_protocol"
)

func main() {

	fmt.Println("Starting echo example.")

	portname := "/dev/ttyUSB0"
	baudrate := 9600

	fmt.Printf("use portname: %s\n", portname)
	fmt.Printf("use baudrate: %d\n", baudrate)

	arduino := simple_serial_protocol.NewSsp(portname, baudrate)

	fmt.Printf("Registering command.\n")

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

	fmt.Println("Connected. Listening for commands.")

	write := config.NewWriteCommandConfig(
		byte('r'),
		[]any{
			byte(0xff),
			true,
			int8(-128),
			uint8(255),
			int16(-32768),
			uint16(65535),
			int32(-2147483648),
			uint32(4294967295),

			int64(-9223372036854775808),
			uint64(18446744073709551615),

			float32(-1.23456789101112),
			byte('J'),

			"text1: Hey, I'm text one!",
			"text2: And I am his brother text two!",
			"text3: Nice!",
		},
	)

	//
	// Waiting for ENTER key to exit
	//
	stopChan := make(chan struct{})
	go func() {
		_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {

		}
		stopChan <- struct{}{}
	}()

	fmt.Println("Give the arduino some time...")
	time.Sleep(3 * time.Second)

	fmt.Println("Writing command.")
	err = arduino.WriteCommand(write)
	if err != nil {
		fmt.Printf("Could not write command: %s\n", err)
		return
	}
	fmt.Println("Command written.")

	fmt.Println("Press ENTER to close application...")
	<-stopChan
	fmt.Println("Bye bye.")
}

func onRead(values []any, err error) {

	if err != nil {
		fmt.Printf("Got an error: %s\n", err)
		return
	}

	fmt.Printf("byteValue: %#x \n", values[0])
	fmt.Printf("booleanValue: %t \n", values[1])
	fmt.Printf("int8Value: %v \n", values[2])
	fmt.Printf("uint8Value: %d \n", values[3])
	fmt.Printf("int16Value: %d \n", values[4])
	fmt.Printf("uint16Value: %d \n", values[5])
	fmt.Printf("int32Value: %d \n", values[6])
	fmt.Printf("uint32Value: %d \n", values[7])
	fmt.Printf("int64Value: %d \n", values[8])
	fmt.Printf("uint64Value: %d \n", values[9])
	fmt.Printf("float32Value: %f \n", values[10])
	fmt.Printf("charValue: %#x \n", values[11])
	fmt.Printf("stringValue1: %s \n", values[12])
	fmt.Printf("stringValue2: %s \n", values[13])
	fmt.Printf("stringValue3: %s \n", values[14])

	fmt.Println("I have successfully received all the data from the Arduino. Bye bye.")

	os.Exit(0)
}
