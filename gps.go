package main

import (
	"log"

	serial "go.bug.st/serial.v1"
)

func InitGPS(portDevice string) (serial.Port, []byte) {

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(portDevice, mode)
	if err != nil {
		log.Fatal(err)
	}

	_, err = port.Write([]byte("AT+CGNSPWR=1\r\nAT+CGNSSEQ=\"RMC\"\r\nAT+CGNSINF\r\nAT+CGNSURC=2\r\nAT+CGNSTST=1\r\n"))

	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 400)

	return port, buff
}
