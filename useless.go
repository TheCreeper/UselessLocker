// Package useless contains all the nessessary code for useless locker.
package useless

import (
	"github.com/TheCreeper/UselessLocker/useless/store"
)

func Start() (err error) {
	s, err := store.Open()
	if err != nil {
		return
	}
	b, err := s.ReadFile("/assets/master.pem")
	if err != nil {
		return
	}
	println(string(b))
	return
}
