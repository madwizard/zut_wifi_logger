package main

import (
	"go.bug.st/serial.v1"
	"log"
	"strconv"
	"time"
	"strings"
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

func writeGpsData(input string) {
	now := time.Now()
	timestamp := now.Unix()

	var tmp []string

	lines := strings.Split(input, "\n")

	// Example line we're lookig for:
	// $GNRMC,211622.000,A,5327.819159,N,01432.634518,E,0.00,152.07,310319,,,A*7B
	// Interesing part is: 5327.819159,N,01432.634518,E
	for _, line := range lines {
		if strings.Contains(line, "GNRMC") {
			if len(line) >= 49 {
				GpsData.Timestamp = strconv.FormatInt(timestamp, 10)
				tokens := strings.Split(line, ",")
				for i := 3; i <= 6; i++ {
					tmp = append(tmp, tokens[i])
				}
				ret := strings.Join(tmp, ",")
				GpsData.Data = ret
				log.Printf("Output: %s", GpsData.Data)
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