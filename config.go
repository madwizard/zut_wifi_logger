package main

import (
	"io/ioutil"
	"github.com/olebedev/config"
	"log"
)

func readConfig(pathname string) (err error) {
	file, err := ioutil.ReadFile(pathname)
	if err != nil {
		log.Printf("Couldn't read config file: %v", err)
		return err
	}
	yamlString := string(file)

	conf, err := config.ParseYaml(yamlString)
	if err != nil {
		log.Printf("Couldn't parse config file: %v", err)
		return  err
	}
	WiFi, err = conf.String("devices.interface")
	if err != nil {
		return err
	}
	usb, err = conf.String("devices.usbport")
	if err != nil {
		return err
	}
	return nil
}
