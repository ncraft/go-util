package http

import (
	"crypto/subtle"
	"net/http"
)

type HandlerWithHttpAuth struct {
	Username        string
	Password        string
	Realm           string
	OriginalHandler http.Handler
}

func (h HandlerWithHttpAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(h.Username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(h.Password)) != 1 {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+h.Realm+`"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))
		return
	}

	h.OriginalHandler.ServeHTTP(w, r)
}
