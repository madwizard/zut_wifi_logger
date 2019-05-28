package main

import (
	"io/ioutil"
	"strings"
)


// setWiFiInterface reads config file and sets interface name
func setWiFiInterface(pathname string) (string, error) {
	f, err := ioutil.ReadFile(pathname)
	if err != nil	{
		return "Error", err
	}
	return strings.TrimRight(string(f), "\n"), nil
}