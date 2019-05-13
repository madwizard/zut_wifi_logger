package main

import (
	"testing"
)



func TestReturnData(t *testing.T) {

}

func TestReadWiFiList(t *testing.T) {
	out, err := readList("wlp4s0")
	if err != "" {
		t.Errorf("readWiFiList returned error '%s'", err)
	}
	if out == "" {
		t.Errorf("readWiFiList returned empty list")
	}
}
