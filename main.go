package main

import (
	"flag"
	"log"
	"fmt"
)

const version string = "1.2.0"

// command line flags
var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Println("timer version", version)
	} else {
		// TODO
	}
}
