package main

import (
	"html/template"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var GpsData gpsData

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	initDB()
}

func main() {
	stopScanner := make(chan bool)
	stopGpsScanner := make(chan bool)

	GpsData.Data = "ABCDEADCOW"
	GpsData.Timestamp = "ABCNEVER"

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
