package simple_serial_protocol

import (
	"go.bug.st/serial"
	"testing"
	"time"
)

//
//
// Useful things
//
// https://github.com/stretchr/testify
// https://github.com/bugst/go-serial/blob/master/serial_linux_test.go
//

func startMockSerialServer(port string) (serial.Port, error) {
	// Öffne den Port
	mode := &serial.Mode{
		BaudRate: 9600,
	}

	s, err := serial.Open(port, mode)
	if err != nil {
		return nil, err
	}

	go func() {
		buf := make([]byte, 128)
		for {
			n, err := s.Read(buf)
			if err != nil {
				return
			}
			_, err = s.Write(buf[:n]) // Echo der empfangenen Daten
			if err != nil {
				return
			}
		}
	}()

	return s, nil
}

func TestSimpleSerialProtocol_Open(t *testing.T) {

	port := "/dev/ttyUSB0" // Beispiel für Windows, oder "/dev/ttyUSB0" für Linux/Mac

	// Starte den simulierten Echo-Server
	server, err := startMockSerialServer(port)
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer server.Close()

	// Konfiguriere den Client für die serielle Kommunikation
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	client, err := serial.Open(port, mode)
	if err != nil {
		t.Fatalf("Failed to open serial port: %v", err)
	}
	defer client.Close()

	// Testdaten senden
	dataToSend := []byte("hello")
	_, err = client.Write(dataToSend)
	if err != nil {
		t.Fatalf("Failed to send data: %v", err)
	}

	// Warte auf das Echo
	time.Sleep(100 * time.Millisecond)

	// Daten empfangen
	buf := make([]byte, 128)
	n, err := client.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read data: %v", err)
	}

	// Überprüfen, ob die gesendeten Daten zurückgekommen sind
	receivedData := buf[:n]
	if string(receivedData) != string(dataToSend) {
		t.Errorf("Expected %s, but got %s", string(dataToSend), string(receivedData))
	}
}
