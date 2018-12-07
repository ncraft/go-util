package main

import "github.com/nextcraft-gmbh/go-util/pkg/log"

func main() {
	log.SetDebug(true)

	log.Info("this is a log example")

	items := []string{"one", "two", "three"}

	for _, color := range items {
		log.Debug("added color %s", color)
	}
}
