package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// Inspired by https://talks.golang.org/2014/testing.slide#23
func TestUsage(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestUsage")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	// capture output of process execution
	r, w, _ := os.Pipe()
	cmd.Stderr = w
	err := cmd.Run()
	w.Close()

	// check return code
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatalf("Expected exit status 1, but was: %v, ", err)
	}

	// now check that Usage message is displayed
	captured, _ := ioutil.ReadAll(r)
	actual := string(captured)
	expected := "Usage:"

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

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
