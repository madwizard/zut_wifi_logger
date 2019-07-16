package main

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

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

// returnData gets a line, splits it on mask and returns second token
func returnData(input string, mask string) string {
	var tmp []string
	if strings.Contains(input, mask) {
		tmp = strings.Split(input, mask)
	}
	return tmp[1]
}

// pack packs data from input to struct
func pack(input string) *wifiData{

	var scannedData wifiData
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Don't capture scan output header nor Unknown fields
		if strings.Contains(line, "IE: Unknown") || strings.Contains(line, "Scan completed") {
			continue
		}
		if strings.Contains(line, "Address: ") {
			scannedData.MAC = returnData(line, "Address: ")
		}
		if strings.Contains(line, "Channel:") {
			scannedData.Channel, _ = strconv.Atoi(returnData(line, "Channel:"))
		}
		if strings.Contains(line, "Frequency:") {
			scannedData.Freq = returnData(line, "Frequency:")
		}
		if strings.Contains(line, "Quality") {
			continue // TBD - Quality=62/70  Signal level=-48 dBm
		}
		if strings.Contains(line, "Encryption") {
			scannedData.Enc = returnData(line, "Encryption key:")
		}
		if strings.Contains(line, "ESSID") {
			scannedData.ESSID = returnData(line, "ESSID:")
		}
	} // End of for
	return &scannedData
}

// readList calls iwlist command and reads in scanned data
func readList(NIC string) (string, string) {

	cmd := exec.Command("/usr/sbin/iwlist", NIC, "scanning")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with error %s and coundn't be run\n", err)
	}

	return string(stdout.String()), string(stderr.String())
}

// parse parses scanned data and packs into slice of wifiData
func WiFiParse(NIC string) *[]wifiData {
	var ret []wifiData
	read, err := readList(NIC)
	if err != "" {
		log.Printf("WiFiParse failed with error %s and couldn't be run\n", err)
		return nil
	}
	readSlice := strings.Split(read, "Cell")

	for _, singleRead := range readSlice {
		w := pack(singleRead)
		ret = append(ret, *w)
	} // End of for

	return &ret
}


func Scanner(stop chan bool) {
	stopscanner := false


	WIFI, err := setWiFiInterface("config.yml")
	if err != nil {
		log.Fatal("Scanner: Can't read config file!")
	}


	log.Println("Scanner: starting")

	for {
		var ScannedData *[]wifiData
		ScannedData = WiFiParse(WIFI)

		now := time.Now()
		timestamp := now.Unix()
		writeWiFiDB(*ScannedData, timestamp)


		select {
			case stopscanner = <- stop:
				if stopscanner {
					log.Println("Scanner: stopping")
					break
				}
				default:
					continue
		}
	}
}