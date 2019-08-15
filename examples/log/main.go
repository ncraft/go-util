package main

import "github.com/ncraft/machinery/pkg/log"

func main() {
	log.SetDebug(true)

	log.Info("this is a log example")

	colors := []string{"red", "blue", "green"}

	for _, color := range colors {
		log.Debug("added color %s", color)
	}
}
