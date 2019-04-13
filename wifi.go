package main

import
(
	"bytes"
	"log"
	"os/exec"
)

func readWifiList(NIC string) (string, string){
	cmd := exec.Command("/usr/sbin/iwlist", NIC, "scanning")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with error %s\n", err)
	}

	return string(stdout.Bytes()), string(stderr.Bytes())

}
/*
func wifiInfo(wifilist string) (string, string, string) {
	temp := strings.Split(wifilist, "\n")
	return
}
 */