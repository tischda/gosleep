package main

import (
	"os"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	params := []string{"0.1", "0.105s", "100ms", "100000us", "100000000ns"}
	args := os.Args

	for _, duration := range params {
		os.Args = append(args, duration)

		start := time.Now()
		main()
		stop := time.Now()

		actual := stop.Sub(start)
		expected := time.Duration(100 * time.Millisecond)

		if actual < expected || actual > (expected+(10*time.Millisecond)) {
			t.Errorf("Duration: %s, Expected: %s, but was: %s", duration, expected, actual)
		}
	}
}

type Test struct {
	in  string
	out bool
}

func TestOnlyNumbers(t *testing.T) {
	cases := []Test{
		{"1h5m", false},
		{"0.5", true},
		{"1", true},
	}
	for _, test := range cases {
		actual := onlyNumbers(test.in)
		if actual != test.out {
			t.Errorf("Expected for %s: %t, but was: %t", test.in, test.out, actual)
		}
	}
}
