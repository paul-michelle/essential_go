package main

import (
	"fmt"
	"log"
	"os"

	"errorHandling/cfg"
)

func setupLogging() {
	fileToLogTo, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(fileToLogTo)
}

func main() {
	setupLogging()

	configs, err := cfg.ReadConfig("fakeConfig")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("configs: %+v\n", configs)
}
