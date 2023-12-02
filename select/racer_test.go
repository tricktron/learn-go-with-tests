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

	slowServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

	fastServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

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
