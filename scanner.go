package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)


func init() {

	err := readConfig("./config.yml")
	if err != nil {
		log.Printf("Couldn't read config file: %v", err)
		os.Exit(-1)
	}

	port = InitGPS(usb)

	initDB()
}

func main() {
	stopScanner := make(chan bool)
	stopGpsScanner := make(chan bool)

	GPSdata.Timestamp = "00000000"
	GPSdata.Latitude = "00000000"
	GPSdata.Longitute = "00000000"
	GPSdata.GPSRead = false

	// Gracefully close down
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go gpsScanner(stopGpsScanner)
	go Scanner(stopScanner)

	sig := <- sigs
	log.Printf("main: Got signal: %+v", sig)
	stopScanner <- true
	stopGpsScanner <- true
	log.Printf("main: Shutting down")
}
