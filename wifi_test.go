package main

import (
	"testing"
)

func TestReadWiFiList(t *testing.T) {
	out, err := readWiFiList("wlp4s0")
	if err != "" {
		t.Errorf("readWiFiList returned error '%s'", err)
	}
	if out == "" {
		t.Errorf("readWiFiList returned empty list")
	}
}

func TestGetKernelWEVersion(t *testing.T) {
	out := getKernelWEVersion()
	expected := 22
	if out != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, out)
	}
}
