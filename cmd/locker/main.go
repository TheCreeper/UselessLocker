package main

import (
	"log"

	"github.com/TheCreeper/UselessLocker/useless"
)

func main() {
	// Safety Pin
	if err := useless.Start(); err != nil {
		log.Fatal(err)
	}
}
