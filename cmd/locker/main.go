package main

import (
	"log"

	"github.com/TheCreeper/UselessLocker/useless"
)

func main() {
	// Safety pin
	if true {
		return
	}

	if err := useless.Start(); err != nil {
		log.Fatal(err)
	}
}
