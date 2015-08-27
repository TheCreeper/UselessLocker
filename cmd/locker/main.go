package main

import (
	"log"

	"github.com/TheCreeper/UselessLocker/useless"
)

func main() {
	if err := useless.Start(); err != nil {
		log.Fatal(err)
	}
}
