package racer_test

import (
	"testing"

	racer "learn-go-with-tests/select"
)

func TestRacer(t *testing.T) {
	t.Parallel()

	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := racer.Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
