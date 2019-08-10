package main

import (
	"html/template"
	"log"
	"os"
	"os/signal"
	"syscall"
	"go.bug.st/serial.v1"
)

var GPSdata GpsData
var tmpl *template.Template
var port serial.Port
var usb string
var WiFi string

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

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

	// Gracefully close down
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go gpsScanner(stopGpsScanner)
	go Scanner(stopScanner)
	go wwwServer()

	sig := <- sigs
	log.Printf("main: Got signal: %+v", sig)
	stopScanner <- true
	stopGpsScanner <- true
	log.Printf("main: Shutting down")
}
