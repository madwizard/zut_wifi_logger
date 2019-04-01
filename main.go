package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("Initialising devices.")

	port, buff := InitGPS("/dev/ttyUSB0")
	defer port.Close()

	fmt.Println("=====================================")
	fmt.Println("Listing WiFis:")
	out, _ := readWifiList("wlp2s0")

	fmt.Println(out)
	fmt.Println("=====================================")
	fmt.Println("Position:")

	for i := 0; i < 10; i++ {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}

		if n == 0 {
			fmt.Println("\nEOF")
		}
		fmt.Println("=====================================")
		fmt.Println(string(buff[:n]))
	}
}
