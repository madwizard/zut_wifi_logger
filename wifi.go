package main

/*
#cgo CFLAGS: -I./cinclude
#cgo LDFLAGS: -L./dynlib/ -lm -liw
#include "iwlib.h"
*/
import "C"

import
(
	"bytes"
	"log"
	"os/exec"
)

func readWiFiList(NIC string) (string, string){
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

func getKernelWEVersion() (int) {
	return int(C.iw_get_kernel_we_version())
}


/*
func wifiInfo(wifilist string) (string, string, string) {
	temp := strings.Split(wifilist, "\n")
	return
}
 */