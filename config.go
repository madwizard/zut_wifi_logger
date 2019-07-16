package main

import (
	"io/ioutil"
	"github.com/olebedev/config"
	"log"
)

type Cfg struct {
	wifiName string
	usbName string
}

func readConfig(pathname string) (cfg *Cfg, err error) {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("Couldn't read config file: %v", err)
		return nil, err
	}
	yamlString := string(file)

	conf, err := config.ParseYaml(yamlString)
	if err != nil {
		log.Printf("Couldn't parse config file: %v", err)
		return nil, err
	}
	cfg.wifiName, err = conf.String("devices.interface")
	if err != nil {
		return nil, err
	}
	cfg.usbName = conf.String("devices.usbport")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
