package main

import (
	"fmt"
	httpUtil "github.com/ncraft/go-util/pkg/http"
	"net/http"
)

type colors []string

func main() {
	colors := colors{"red", "blue", "green"}

	authHandler := httpUtil.HandlerWithHttpAuth{
		Username:        "user123",
		Password:        "secret",
		Realm:           "Please enter your username and password for this site",
		OriginalHandler: http.HandlerFunc(colors.list),
	}

	http.Handle("/colors", authHandler)

	http.ListenAndServe(":8080", nil)
}

func (colors colors) list(w http.ResponseWriter, req *http.Request) {
	for _, c := range colors {
		fmt.Fprintf(w, "%s\n", c)
	}
}
