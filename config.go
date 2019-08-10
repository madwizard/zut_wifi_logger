package main

import (
	"go.bug.st/serial.v1"
	"io/ioutil"
	"github.com/olebedev/config"
	"log"
)

var port serial.Port
var usb string
var WiFi string
var GPSdata GpsData

type GpsData struct {
	Timestamp string
	Latitude string
	Longitute string
}

// Example scanned data is in docs/example_WiFiScanData
type wifiData struct {
	ESSID string `json:"ESSID"` 		// ESSID
	MAC string	`json:"MAC"`			// Address
	Freq string `json:"freq"`			// Frequency
	SigLvl string `json:"siglvl"`		// SignalLevel
	Qual string `json:"qual"`			// Quality
	Enc string `json:"enc"`				// Encryption key
	Channel int `json:"channel"`		// Channel
	Mode string `json:"mode"`			// Mode
	IEEE string `json:"IEEE"`			// IEEE
	Bitrates string `json:"bitrates"`	// bitrates
	WPA string `json:"wpa"`				// WPA version
}

type webdata struct {
	Timestamp string `json:"Timestamp"`
	ESSID string `json:"ESSID"` 		// ESSID
	MAC string	`json:"MAC"`			// Address
	Freq string `json:"freq"`			// Frequency
	SigLvl string `json:"siglvl"`		// SignalLevel
	Qual string `json:"qual"`			// Quality
	Enc string `json:"enc"`				// Encryption key
	Channel int `json:"channel"`		// Channel
	Mode string `json:"mode"`			// Mode
	IEEE string `json:"IEEE"`			// IEEE
	Bitrates string `json:"bitrates"`	// bitrates
	WPA string `json:"wpa"`				// WPA version
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

func readConfig(pathname string) (err error) {
	file, err := ioutil.ReadFile(pathname)
	if err != nil {
		log.Printf("Couldn't read config file: %v", err)
		return err
	}
	yamlString := string(file)

	conf, err := config.ParseYaml(yamlString)
	if err != nil {
		log.Printf("Couldn't parse config file: %v", err)
		return  err
	}
	WiFi, err = conf.String("devices.interface")
	if err != nil {
		return err
	}
	usb, err = conf.String("devices.usbport")
	if err != nil {
		return err
	}
	return nil
}
