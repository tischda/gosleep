package main

import (
	"flag"
	"log"
	"fmt"
	"time"
	"os"
	"regexp"
)

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
// go build -ldflags "-x main.version $(git describe --tags)"
var version string

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
		delay := flag.Arg(0)
		if onlyNumbers(delay) {
			delay += "s" // seconds
		}
		d, err := time.ParseDuration(delay)
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

func onlyNumbers(value string) bool {
	exp := regexp.MustCompile(`^[\d.]+$`)
	return exp.FindString(value) != ""
}
