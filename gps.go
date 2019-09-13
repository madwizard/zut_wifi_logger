package main

import (
	"github.com/adrianmo/go-nmea"
	"go.bug.st/serial.v1"
	"log"
	"strconv"
	"strings"
	"time"
)

func InitGPS(portDevice string) (serial.Port) {

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("/dev/" + portDevice, mode)
	if err != nil {
		log.Printf("%v", err)
		startGPS = false
		return nil
	}

	startGPS = true
	
	_, err = port.Write([]byte("AT+CGNSPWR=1\r\nAT+CGNSSEQ=\"RMC\"\r\nAT+CGNSINF\r\nAT+CGNSURC=2\r\nAT+CGNSTST=1\r\n"))

	if err != nil {
		log.Fatal(err)
	}

	return port
}

func ReadGPS(port serial.Port) string {
	buff := make([]byte, 1)
	var ret strings.Builder

	for string(buff[0]) != "\n" {
	_, err := port.Read(buff)
		if err != nil {
			log.Printf("Couldn't read GPS coords")
		}
		log.Printf("%v", string(buff[0]))
		ret.WriteString(string(buff[0]))
	}
	log.Printf("GPS DATA: %v", ret.String())
	return ret.String()
}

func writeGpsData(input string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if strings.Contains(line, "GNRMC") {
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
					GPSdata.Latitude = convertDMStoDec(nmea.FormatGPS(m.Latitude))
					GPSdata.Longitute = convertDMStoDec(nmea.FormatGPS(m.Longitude))
					GPSdata.GPSRead = true
				}
		} else {
			continue
		}
	}
}

func convertDMStoDec(data string) string {
	var ret strings.Builder
	tmp := strings.Split(data, ".")
	tmp1 := tmp[0]
	secDec, _ := strconv.Atoi(tmp[1])						// Seconds Decimal
	minDec, _ := strconv.Atoi(tmp1[len(tmp1)-2:])			// Minutes Decimal
	degDec, _ := strconv.Atoi(tmp1[:len(tmp1)-2])			// Seconds Decimal

	minDec *= 100
	minDec = minDec / 60
	minDec = minDec + secDec

	ret.WriteString(strconv.Itoa(degDec) + "." + strconv.Itoa(minDec))
	// ret.WriteString(degDec + "." + strconv.Itoa(minutes) + secDec)
	return ret.String()
}

func gpsScanner(stop chan bool) {
	stopscanner := false

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
