package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

// number of progress bar increments
var TICKS = len(bar)

func sleep(duration string) {
	// only numbers or dot, unit missing?
	exp := regexp.MustCompile(`^[\d.]+$`)
	if exp.FindString(duration) != "" {
		duration += "s" // seconds
	}
	d, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
	}
	if !quiet {
		go show_progress(d)
	}
	time.Sleep(d)
	if !quiet {
		// no go routine here to make sure the '100%' is printed
		fmt.Println(bar[TICKS-1])
	}
}

// prints the progress bar as time is passing by
func show_progress(d time.Duration) {
	interval := d / time.Duration(TICKS)
	hide_cursor()
	for i := 0; i < TICKS; i++ {
		fmt.Print(bar[i])
		time.Sleep(interval)
	}
	show_cursor()
}
