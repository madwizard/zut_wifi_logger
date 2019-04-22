package main

/*
#cgo CFLAGS: -I./cinclude
#cgo LDFLAGS: -L./dynlib/ -lm -liw
#include "iwlib.h"
*/
import "C"

type wifiData struct {
	essid string
	mac string
	freq string
	siglvl int8
	qual int16
	enc bool
	bitrates []int16
}

func readWiFiList(NIC string) (string, string){
/*	p := C.malloc(C.wireless_scan_head)
	defer C.free(p)

	cmd := exec.Command("/usr/sbin/iwlist", NIC, "scanning")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with error %s\n", err)
	}

	return string(stdout.String()), string(stderr.String())
*/


return "", ""
}

func wirelessScan() int {

	/*
	p := C.malloc(unsafe.Pointer(C.wireless_scan_head{}))
	defer C.free(p)

	iface := C.CString("wlp4s0")
	sk := iwSocketsOpen()
	weVersion := getKernelWEVersion()

	ret := int(C.iw_process_scan(_Ctype_int(sk), iface, _Ctype_int(weVersion), unsafe.Pointer(p)))
 	*/
	ret := 1
	return ret
}

func getKernelWEVersion() (int) {
	//return int(C.iw_get_kernel_we_version())
	return 1
}

func iwSocketsOpen() int {
	// return int(C.iw_sockets_open())
	return 1
}

func iwSocketsClose(skfd int) {
	// C.iw_sockets_close(_Ctype_int(skfd))
}
