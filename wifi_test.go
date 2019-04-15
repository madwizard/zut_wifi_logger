package main

import (
	"testing"
)

func TestReadWiFiList(t *testing.T) {
	out, _ := readWiFiList("wlp2s0")
	if out == "" {
		t.Error("readWifiList returned empty list")
	}
}

func TestgetKernelWEVersion(t *testing.T) {
	out := getKernelWEVersion()
	expected := 22
	if out != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, out)
	}
}
