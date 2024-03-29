package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	config := getConfig()

	startServer(config)
}

func getConfig() Config {
	var errorMsg error
	var port, strMin, strMax int

	port, errorMsg = strconv.Atoi(os.Getenv("PASSMAN_PORT"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}
	strMin, errorMsg = strconv.Atoi(os.Getenv("PASSMAN_STRING_MIN"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}
	strMax, errorMsg = strconv.Atoi(os.Getenv("PASSMAN_STRING_MAX"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}

	c := Config{
		Port:            port,
		StringMinLength: strMin,
		StringMaxLength: strMax,
	}

	return c
}

type Config struct {
	Port            int
	StringMaxLength int
	StringMinLength int
}
