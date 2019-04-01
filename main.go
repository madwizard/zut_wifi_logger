package main

import (
	"fmt"
	"log"
)

func main() {
	port, buff := InitGPS("/dev/ttyUSB0")
	defer port.Close()

	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}

		if n == 0 {
			fmt.Println("\nEOF")
		}
		fmt.Println("%v", string(buff[:n]))
	}
}
