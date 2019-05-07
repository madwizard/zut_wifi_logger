package main

import (
	"bytes"
	"log"
	"os/exec"
)

type wifiData struct {
	ESSID string `json:"ESSID"`
	MAC string	`json:"MAC"`
	Freq string `json:"freq"`
	SigLvl int8 `json:"siglvl"`
	Qual int16 `json:"qual"`
	Enc bool `json:"enc"`
	BitRates []int16 `json:"bitrates"`
}

func (w *wifiData) readWiFiList(NIC string) (string, string){

	cmd := exec.Command("/usr/sbin/iwlist", NIC, "scanning")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with error %s\n", err)
	}

	return string(stdout.String()), string(stderr.String())

}

func (w *wifiData) parseWiFiData(NIC string) {
	_, err := w.readWiFiList(NIC)
	if err != "" {
		log.Fatalf("readWiFiList failed with error %s\n", err)
	}


}