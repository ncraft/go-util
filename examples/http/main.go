package main

import (
	"fmt"
	"net/http"

	httpUtil "github.com/ncraft/go-util/pkg/http"
)

type colors []string

func main() {
	colors := colors{"red", "blue", "green"}

	authHandler := httpUtil.BasicAuthHandler{
		Username:        "user123",
		Password:        "secret",
		Realm:           "Please provide username and password",
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
