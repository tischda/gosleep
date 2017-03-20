package main

import (
	"flag"
	"fmt"
	"os"
)

var version string
var quiet bool

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: gosleep [options] duration (default unit = seconds)")
		fmt.Fprintln(os.Stderr, "\nOPTIONS:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	var showVersion = flag.Bool("version", false, "print version and exit")
	flag.BoolVar(&quiet, "quiet", false, "do not print anything")

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
	}
	if flag.Arg(0) == "version" || *showVersion {
		fmt.Println("gosleep version", version)
	} else {
		sleep(flag.Arg(0))
	}
}
