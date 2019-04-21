package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page")
	fmt.Fprint(w, "Home page")
}

func startScanning(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting scanning")
	fmt.Fprint(w, "Starting scan\n")
	fmt.Fprint(w, "Scanning")
}

func sendData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sending data")
	fmt.Fprint(w, "Sending data")
}