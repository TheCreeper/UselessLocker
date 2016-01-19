// Package verbs provides a simple method to filter request method with the
// net/http ServeMux.
package verbs

import "net/http"

type Verbs struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Put    http.HandlerFunc
	Delete http.HandlerFunc
	Patch  http.HandlerFunc
	Head   http.HandlerFunc
}

func (v Verbs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if v.Get != nil {
			v.Get(w, r)
			return
		}
	case "POST":
		if v.Post != nil {
			v.Post(w, r)
			return
		}
	case "Put":
		if v.Put != nil {
			v.Put(w, r)
			return
		}
	case "Delete":
		if v.Delete != nil {
			v.Delete(w, r)
			return
		}
	case "Patch":
		if v.Patch != nil {
			v.Patch(w, r)
			return
		}
	case "HEAD":
		if v.Head != nil {
			v.Head(w, r)
			return
		}
	}

	http.Error(w,
		http.StatusText(http.StatusMethodNotAllowed),
		http.StatusMethodNotAllowed)
}
