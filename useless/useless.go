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

	/*cfgBytes, err := s.Load(PathStoreConfig)
	if err != nil {
		return
	}

	cfg, err := config.ParseBytes(cfgBytes)
	if err != nil {
		return
	}*/

	key, err := CreateSession(s)
	if err != nil {
		return
	}

	if err = EncryptHome(key); err != nil {
		return
	}
	return
}

func OpenBrowser(s store.Store) (err error) {
	//return browser.OpenFile(filename)
	return
}
