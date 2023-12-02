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
	fastServer := makeDelayedServer(0 * time.Millisecond)
	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL

	got := racer.Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
