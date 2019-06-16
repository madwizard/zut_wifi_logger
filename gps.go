package main

import (
	"go.bug.st/serial.v1"
	"log"
	"time"
)

type gpsData struct {
	Data string
	Timestamp string
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

func gpsScanner(stop chan bool) {
	stopscanner := false

	port := InitGPS("/dev/ttyUSB0")

	log.Println("Scanner: starting")

	for {

		log.Printf("Scanner: pass")
		data := ReadGPS(port)

		now := time.Now()
		timestamp := now.Unix()
		writeGpsDB(data, timestamp)


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