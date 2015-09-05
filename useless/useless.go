// Package useless contains all the nessessary code for useless locker.
package useless

import (
	//"github.com/TheCreeper/UselessLocker/useless/browser"
	//"github.com/TheCreeper/UselessLocker/useless/config"
	//"github.com/TheCreeper/UselessLocker/useless/home"
	"github.com/TheCreeper/UselessLocker/useless/store"
)

func Start() (err error) {
	s, err := store.Open()
	if err != nil {
		return
	}
	defer s.Close()

	pubBytes, err := s.Load(PathPublicKey)
	if err != nil {
		return
	}

	key, err := CreateSession(pubBytes)
	if err != nil {
		return
	}

	if err = EncryptHome(key); err != nil {
		return
	}
	return
}
