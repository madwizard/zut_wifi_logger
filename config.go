package main

import (
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Config struct {
	wifiName string `json:"wifiName"`
	usbName string `json:"usbName"`

}

func readConfig(pathname string) error {

}
