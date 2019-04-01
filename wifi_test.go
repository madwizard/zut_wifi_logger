package main

import (
	"testing"
)

func TestWifi(t *testing.T) {
	out, _ := readWifiList("wlp2s0")
	if out == "" {
		t.Error("readWifiList returned empty list")
	}
}
