[![Build Status](https://travis-ci.com/ncraft/machinery.svg?branch=master)](https://travis-ci.com/ncraft/machinery) [![GoDoc](https://godoc.org/github.com/ncraft/machinery?status.svg)](http://godoc.org/github.com/ncraft/machinery)

# Machinery

Machinery for logging, configuration, http basic auth etc.

## HTTP Basic authentication

Example:
```go
package main

import (
	"fmt"
	"net/http"
	
	httpUtil "github.com/ncraft/machinery/pkg/http"
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
```

## Logging

Example:
```go
package main

import "github.com/ncraft/machinery/pkg/log"

func main() {
	log.SetDebug(true)

	log.Info("running logging example")

	colors := []string{"red", "blue", "green"}

	for _, color := range colors {
		log.Debug("added color %s", color)
	}
}
```

Prints:
```
INFO: 2018/12/07 11:28:23 main.go:8: this is a log example
DEBUG: 2018/12/07 11:28:23 main.go:13: added color one
DEBUG: 2018/12/07 11:28:23 main.go:13: added color two
DEBUG: 2018/12/07 11:28:23 main.go:13: added color three
```
