package main

import (
	"log"
	"testing"
	"time"
)

// acceptable delta for the time effectively slept
const DELTA_MS = 20 * time.Millisecond

func TestSleep(t *testing.T) {
	quiet = true
	params := []string{"0.1", "0.100s", "100ms", "100000us", "100000000ns"}

	for _, duration := range params {
		start := time.Now()
		sleep(duration)
		stop := time.Now()

		actual := stop.Sub(start)
		expected, err := time.ParseDuration("0.1s")
		if err != nil {
			log.Fatalln(err)
		}
		if actual < expected-DELTA_MS || actual > expected+DELTA_MS {
			t.Errorf("Duration: %s, Expected: %s, but was: %s", duration, expected, actual)
		}
	}
}
