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
	for i := 0; i < 100; i++ {
		if os.Getenv("BE_CRASHER") == "1" {
			main()
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestSleep", "0.1")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")

		start := time.Now()
		err := cmd.Run()
		stop := time.Now()

		// check return code
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			t.Fatalf("Expected exit status 0, but was: %v, ", err)
		}

		actual := stop.Sub(start)
		expected := time.Duration(100 * time.Millisecond)

		if actual <  expected {
			t.Errorf("Expected: >= %s, but was: %s", expected, actual)
		}
	}
}
