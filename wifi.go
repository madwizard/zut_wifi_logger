package main

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
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

func returnData(input string, mask string) string {
	var tmp []string
	if strings.Contains(input, mask) {
		tmp = strings.Split(input, mask)
	}
	return tmp[1]
}
// pack packs data from input to struct
func (w *wifiData) pack(input string) {

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Don't capture scan output header nor Unknown fields
		if strings.Contains(line, "IE: Unknown") || strings.Contains(line, "Scan completed") {
			continue
		}
		if strings.Contains(line, "Address: ") {
			w.MAC = returnData(line, "Address: ")
		}
		if strings.Contains(line, "Channel:") {
			w.Channel, _ = strconv.Atoi(returnData(line, "Channel:"))
		}
		if strings.Contains(line, "Frequency:") {
			w.Freq = returnData(line, "Frequency:")
		}
		if strings.Contains(line, "Quality") {
			continue // TBD - Quality=62/70  Signal level=-48 dBm
		}
		if strings.Contains(line, "Encryption") {
			w.Enc = returnData(line, "Encryption key:")
		}
		if strings.Contains(line, "ESSID") {
			w.ESSID = returnData(line, "ESSID:")
		}
	} // End of for
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
func WiFiParse(NIC string, w* wifiData)  {
	read, err := readList(NIC)
	if err != "" {
		log.Printf("WiFiParse failed with error %s and couldn't be run\n", err)
		return
	}
	readSlice := strings.Split(read, "Cell")

	for _, singleRead := range readSlice {
		w.pack(singleRead)
	} // End of for

}