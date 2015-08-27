package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct{ HTTPServer string }

func Load(filename string) (cfg Config, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &cfg); err != nil {
		return
	}
	return
}
