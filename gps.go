package main

import (
	"go.bug.st/serial.v1"
	"log"
)

type gpsData struct {
	longitude string
	latitude string
}

func InitGPS(portDevice string) (serial.Port) {

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

	return port
}

func ReadGPS(port serial.Port) string {
	buff := make([]byte, 200)
	n, err := port.Read(buff)
	if err != nil {
		log.Printf("Couldn't read GPS coords")
	}
	return string(buff[:n])
}

func GpsScanner(stop chan bool) {
	stopscanner := false

	port := InitGPS("/dev/ttyUSB0")

	for {
		data := ReadGPS(port)
		writeGpsDB(data)

		select {
		case stopscanner = <- stop:
			if stopscanner == true {
				log.Println("GPS Scanner: stopping")
				port.Close()
				break
			}
		default:
			continue
		}
	}
}