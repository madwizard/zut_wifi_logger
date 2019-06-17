package main

import (
	"go.bug.st/serial.v1"
	"log"
	"strings"
	"github.com/adrianmo/go-nmea"
)

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

func writeGpsData(input string) {
	lines := strings.Split(input, "\n")

	// Example line we're lookig for:
	// $GNRMC,211622.000,A,5327.819159,N,01432.634518,E,0.00,152.07,310319,,,A*7B
	// $GNRMC,192254.000,A,5327.807738,N,01432.629884,E,0.26,176.34,170619,,,A*7F
	// Interesing part is: 5327.819159,N,01432.634518,E
	for _, line := range lines {
		if strings.Contains(line, "GNRMC") {
			log.Printf("%v", line)
			if len(line) >= 74 {
				s, err := nmea.Parse(line)
				if err != nil {
					log.Printf("Couldn't parse GPS data: %v", err)
					continue
				}
				if s.DataType() == nmea.TypeRMC {
					m := s.(nmea.RMC)
					log.Printf("Raw sentence: %v\n", m)
					log.Printf("Time: %s\n", m.Time)
					log.Printf("Validity: %s\n", m.Validity)
					log.Printf("Latitidue GPS: %s\n", nmea.FormatGPS(m.Latitude))
					log.Printf("Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
					log.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
					log.Printf("Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))
					log.Printf("Speed: %f\n", m.Speed)
					log.Printf("Course: %f\n", m.Course)
					log.Printf("Date: %s\n", m.Date)
					log.Printf("Variation: %f\n", m.Variation)
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