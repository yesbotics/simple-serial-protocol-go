module ssp-example/send

go 1.22

replace yesbotics/ssp => ../..

replace yesbotics/ssp/pkg/simple_serial_protocol => ../../pkg/ssp

require yesbotics/ssp v0.0.0-00010101000000-000000000000

require (
	github.com/creack/goselect v0.1.2 // indirect
	go.bug.st/serial v1.6.2 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
)
