package racer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	racer "learn-go-with-tests/select"
)

func TestRacer(t *testing.T) {
	t.Parallel()

	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL

	got := racer.Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
