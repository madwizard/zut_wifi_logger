package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Example scanned data is in docs/example_WiFiScanData
type wifiData struct {
	ESSID string `json:"ESSID"` 	// ESSID
	MAC string	`json:"MAC"`		// Address
	Freq string `json:"freq"`		// Frequency
	SigLvl string `json:"siglvl"`	// Singla Level
	Qual int16 `json:"qual"`		//
	Enc bool `json:"enc"`
	BitRates []int16 `json:"bitrates"`
	Channel int8 `json:"channel"`

}

// pack packs data from input to struct
func (w *wifiData) pack(input string)  {

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Don't capture scan output header nor Unknown fields
		if strings.Contains(line, "IE: Unknown") || strings.Contains(line, "Scan completed") {
			continue
		}
		fmt.Println(line)
	}
}

// readList calls iwlist command and reads in scanned data
func readList(NIC string) (string, string) {

	cmd := exec.Command("/usr/sbin/iwlist", NIC, "scanning")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with error %s\n", err)
	}

	return string(stdout.String()), string(stderr.String())
}

// parse parses scanned data and packs into slice of wifiData
func WiFiParse(NIC string, w* wifiData)  {
	read, err := readList(NIC)
	if err != "" {
		log.Printf("WiFiParse failed with error %s\n", err)
		return
	}
	readSlice := strings.Split(read, "Cell")

	for _, singleRead := range readSlice {
		w.pack(singleRead)
	} // End of for

}