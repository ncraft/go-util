// HTTP utils for basic authentication etc.
package http

import (
	"crypto/subtle"
	"net/http"
)

// Wraps an original http.Handler with credentials to do http basic authentication and a realm text to be displayed in 401 responses.
//
// Usage:
//  authHandler := httpUtil.BasicAuthHandler{
//      Username: "user123",
//      Password: "secret",
//      Realm:    "Please provide username and password",
//      OriginalHandler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//          // ..
//      }),
//  }
//
//  http.Handle("/sample", authHandler)
type BasicAuthHandler struct {
	Username        string
	Password        string
	Realm           string
	OriginalHandler http.Handler
}

// BasicAuthHandler satisfies http.Handler interface and after checking the credentials delegates to the original handler.
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
