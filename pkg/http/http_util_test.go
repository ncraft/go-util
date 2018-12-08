package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/colors", nil)
	if err != nil {
		t.Fatal(err)
	}

	authHandler := HandlerWithHttpAuth{
		Username: "user123",
		Password: "secret",
		Realm:    "Please enter your username and password for this site",
		OriginalHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

			io.WriteString(w, `{"very_secret_number": 1234}`)
		}),
	}

	// use ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	recorder := httptest.NewRecorder()

	// handler satisfies http.Handler interface, so we can call ServeHTTP method directly and pass in ResponseRecorder
	authHandler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}
