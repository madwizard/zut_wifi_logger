package main

import (
	"go.bug.st/serial.v1"
	"log"
	"strconv"
	"strings"
	"github.com/adrianmo/go-nmea"
	"time"
)

type GpsData struct {
	Timestamp string
	Latitude string
	Longitute string
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
	buff := make([]byte, 300)
	n, err := port.Read(buff)
	if err != nil {
		log.Printf("Couldn't read GPS coords")
	}
	return string(buff[:n])
}

func writeGpsData(input string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if strings.Contains(line, "GNRMC") {
			if len(line) >= 74 {
				line = strings.TrimSuffix(line, "\n")
				s, err := nmea.Parse(line)
				if err != nil {
					log.Printf("Couldn't parse GPS data: %v", err)
					continue
				}
				if s.DataType() == nmea.TypeRMC {
					m := s.(nmea.RMC)
					now := time.Now()
					GPSdata.Timestamp = strconv.FormatInt(now.Unix(), 10)
					GPSdata.Latitude = nmea.FormatGPS(m.Latitude)
					GPSdata.Longitute = nmea.FormatGPS(m.Longitude)
				}
			}
		} else {
			continue
		}
	}
}

func gpsScanner(stop chan bool) {
	stopscanner := false

	port := InitGPS("/dev/ttyUSB0")

	log.Println("Scanner: starting")

	for {
		data := ReadGPS(port)

		writeGpsData(data)

		select {
		case stopscanner = <- stop:
			if stopscanner {
				log.Println("GPS Scanner: stopping")
				port.Close()
				break
			}
		default:
			continue
		}
	}
}