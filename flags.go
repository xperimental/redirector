package main

import (
	"errors"
	"flag"
	"os"
)

var addr string
var destination string
var verbose bool

func init() {
	flag.StringVar(&addr, "addr", ":8080", "Address to bind socket to.")
	flag.StringVar(&destination, "destination", "", "Destination to redirect to.")
	flag.BoolVar(&verbose, "verbose", false, "Log every request.")
}

func checkArgs() error {
	flag.Parse()

	portStr := os.Getenv("PORT")
	if len(portStr) > 0 {
		addr = ":" + portStr
	}

	if len(addr) == 0 {
		return errors.New("Need to provide an address.")
	}

	destinationEnv := os.Getenv("DESTINATION")
	if len(destinationEnv) > 0 {
		destination = destinationEnv
	}

	if len(destination) == 0 {
		return errors.New("Need to provide a destination.")
	}

	return nil
}
