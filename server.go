package main

import (
	"fmt"
	"net/http"
	"time"
)

type scanData struct {
	wifiData
	gpsData
	timestamp time.Time
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page")
	fmt.Fprint(w, "Home page")
}

func startScanning(w http.ResponseWriter, r *http.Request, sendData chan <- string, stopRequest <- chan bool) {
	fmt.Println("Starting scanning")
	fmt.Fprint(w, "Starting scan\n")
	fmt.Fprint(w, "Scanning")
	go wirelessScan()
}

func sendData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sending data")
	fmt.Fprint(w, "SSID: biuro\n")
	fmt.Fprint(w, "Signal strength: 20\n")
	fmt.Fprint(w, "GPS: ...")
}