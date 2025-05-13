package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nikhil")

	got := buffer.String()
	want := "Hello, Nikhil"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
