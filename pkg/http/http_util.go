// Provides some http utils e.g. for basic authentication.
package http

import (
	"crypto/subtle"
	"net/http"
)

// Wraps an original http.Handler with credentials and a realm text to be displayed in 401 responses.
type BasicAuthHandler struct {
	Username        string
	Password        string
	Realm           string
	OriginalHandler http.Handler
}

// BasicAuthHandler satisfies http.Handler interfaces.
func (h BasicAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(h.Username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(h.Password)) != 1 {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+h.Realm+`"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))
		return
	}

	h.OriginalHandler.ServeHTTP(w, r)
}
