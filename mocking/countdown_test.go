package mocking_test

import (
	"bytes"
	"testing"

	"learn-go-with-tests/mocking"
)

func TestCountdown(t *testing.T) {
	t.Parallel()

	buffer := &bytes.Buffer{}

	mocking.Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
