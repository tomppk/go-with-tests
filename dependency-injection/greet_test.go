package main

import (
	"bytes"
	"testing"
)

func TestGreet( t *testing.T) {
	// The Buffer type from "bytes" package implements Writer interface
	// We will use it in our test to send in as our Writer and the we can check
	// what was written to it after we invoke Greet
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}