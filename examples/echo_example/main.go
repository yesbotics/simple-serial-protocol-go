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

	//b := make([]byte, 8)
	//binary.LittleEndian.PutUint64(b, 18446744073709551615)
	//fmt.Printf("byteValue: %#x \n", b)

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

	stopChan := make(chan struct{})
	go func() {
		_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {

		}
		stopChan <- struct{}{}
	}()

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

func onRead(args ...any) {

	fmt.Println("Received several values from Arduino:")
	fmt.Printf("byteValue: %#x \n", args[0])
	fmt.Printf("booleanValue: %t \n", args[1])
	fmt.Printf("int8Value: %d \n", args[2])
	fmt.Printf("uint8Value: %d \n", args[3])
	fmt.Printf("int16Value: %d \n", args[4])
	fmt.Printf("uint16Value: %d \n", args[5])
	fmt.Printf("int32Value: %d \n", args[6])
	fmt.Printf("uint32Value: %d \n", args[7])
	fmt.Printf("int64Value: %d \n", args[8])
	fmt.Printf("uint64Value: %d \n", args[9])
	fmt.Printf("float32Value: %f \n", args[10])
	fmt.Printf("charValue: %#x \n", args[11])
	fmt.Printf("stringValue1: %s \n", args[12])
	fmt.Printf("stringValue2: %s \n", args[13])
	fmt.Printf("stringValue3: %s \n", args[14])

	//fmt.Printf("Read data: %s", args)
}

//#72#ff
//#ff
//#01
//#80
//#ff
//#00#80
//#ff#ff
//#00#00#00#80
//#ff#ff#ff#ff
//#00#00#00#00#00#00#00#80
//#ff#ff#ff#ff#ff#ff#ff#ff
//#ff#ff#ff#ff
//#4a
//#74#65#78#74#31#3a#20#48#65#79#2c#20#49#27#6d#20#74#65#78#74#20#6f#6e#65#21#00
//#74#65#78#74#32#3a#20#41#6e#64#20#49#20#61#6d#20#68#69#73#20#62#72#6f#74#68#65#72#20#74#65#78#74#20#74#77#6f#21#00
//#74#65#78#74#33#3a#20#4e#69#63#65#21#00
//#0a
//
//
//#72#ff#ff#01#80#ff#00#80#ff#ff#00#00#00#80#ff#ff#ff#ff#00#00#00#00#00#00#00#80#ff#ff#ff#ff#ff#ff#ff#ff#ff#ff#ff#ff#4a#74#65#78#74#31#3a#20#48#65#79#2c#20#49#27#6d#20#74#65#78#74#20#6f#6e#65#21#00#74#65#78#74#32#3a#20#41#6e#64#20#49#20#61#6d#20#68#69#73#20#62#72#6f#74#68#65#72#20#74#65#78#74#20#74#77#6f#21#00#74#65#78#74#33#3a#20#4e#69#63#65#21#00#0a
