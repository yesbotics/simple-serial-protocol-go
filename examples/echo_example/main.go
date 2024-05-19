package main

import (
	"bufio"
	"fmt"
	"github.com/yesbotics/simple-serial-protocol-go/pkg/config"
	"github.com/yesbotics/simple-serial-protocol-go/pkg/simple_serial_protocol"
	"os"
	"time"
)

func main() {

	// Press ENTER key to exit
	stopChan := make(chan struct{})
	go func() {
		_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {

		}
		stopChan <- struct{}{}
	}()

	fmt.Println("Starting echo example.")

	// Change this to your partner device's baudrate
	portname := "/dev/ttyUSB0"

	// Change this to your serial port name
	baudrate := 9600

	fmt.Printf("Use portname: %s\n", portname)
	fmt.Printf("Use baudrate: %d\n", baudrate)

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

	fmt.Println("Give the Arduino some time...")
	time.Sleep(3 * time.Second)

	fmt.Println("Writing command.")
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
			float32(-1.2345678),
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
	fmt.Println("Command written.")

	fmt.Println("Press ENTER to close application...")
	<-stopChan
	fmt.Println("Bye bye.")
}

func onRead(params []any, err error) {

	if err != nil {
		fmt.Printf("Got an error: %s\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("Got data from Arduino:")
	fmt.Println("")

	fmt.Printf("byteValue: %#x \n", params[0])
	fmt.Printf("booleanValue: %t \n", params[1])
	fmt.Printf("int8Value: %v \n", params[2])
	fmt.Printf("uint8Value: %d \n", params[3])
	fmt.Printf("int16Value: %d \n", params[4])
	fmt.Printf("uint16Value: %d \n", params[5])
	fmt.Printf("int32Value: %d \n", params[6])
	fmt.Printf("uint32Value: %d \n", params[7])
	fmt.Printf("int64Value: %d \n", params[8])
	fmt.Printf("uint64Value: %d \n", params[9])
	fmt.Printf("float32Value: %f \n", params[10])
	fmt.Printf("charValue: %#x / %s\n", params[11], string(params[11].(byte)))
	fmt.Printf("stringValue1: %s \n", params[12])
	fmt.Printf("stringValue2: %s \n", params[13])
	fmt.Printf("stringValue3: %s \n", params[14])

	fmt.Println("")

	fmt.Println("I have successfully received all the data from the Arduino. Press ENTER to exit.")
}
