package main

import (
	"flag"

	"github.com/TheCreeper/UselessLocker/useless"
)

var (
	Decrypt bool
	Encrypt bool
)

func init() {
	flag.BoolVar(&Decrypt, "d", false, ".")
	flag.BoolVar(&Encrypt, "e", false, ".")
	flag.Parse()
}

func main() {
}
