# go-util

Collection of utils for logging, configuration, http basic auth etc.

## Logging

Example:
```golang
package main

import "github.com/ncraft/go-util/pkg/log"

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
INFO:2018/12/07 11:28:23 main.go:8: this is a log example
DEBUG:2018/12/07 11:28:23 main.go:13: added color one
DEBUG:2018/12/07 11:28:23 main.go:13: added color two
DEBUG:2018/12/07 11:28:23 main.go:13: added color three
```