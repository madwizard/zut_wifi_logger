package main

import (
	"fmt"
	"log"

	"go.bug.st/serial.v1"
)

func main() {

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	_, err = port.Write([]byte("AT+CGNSPWR=1\r\nAT+CGNSSEQ=\"RMC\"\r\nAT+CGNSINF\r\nAT+CGNSURC=2\r\nAT+CGNSTST=1\r\n"))

	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 400)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}

		if n == 0 {
			fmt.Println("\nEOF")
		}
		fmt.Println("%v", string(buff[:n]))
	}
}
