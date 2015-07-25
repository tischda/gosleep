package main

import (
	"flag"
	"log"
	"fmt"
	"time"
	"os"
)

const version string = "1.0.1"

// command line flags
var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	setCustomUsage()

	if showVersion {
		fmt.Println("sleep version", version)
	} else {
		if flag.NArg() != 1 {
			flag.Usage()
		}
		d, err := time.ParseDuration(flag.Arg(0)+"s")
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(d)
	}
}

// Redefines flag.Usage() function.
func setCustomUsage() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s seconds\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}