package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Parallel()

	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
