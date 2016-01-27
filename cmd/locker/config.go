package main

import (
	"encoding/json"
)

type Config struct {
	// Filepath to the encrypted session key file which is used in case
	// the key is lost.
	SessionKeyFile string

	// Remote server to push encrypted key to.
	HTTPServer string

	// Max size in megabytes of files to look for to be encrypted.
	MaxFileSize int64
}

func ParseConfig(b []byte) (cfg Config, err error) {
	err = json.Unmarshal(b, &cfg)
	return
}
