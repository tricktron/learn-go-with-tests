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

	t.Run("returns the faster server url", func(t *testing.T) {
		t.Parallel()
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL

		got, _ := racer.Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server does not respond within 10s",
		func(t *testing.T) {
			t.Parallel()
			serverA := makeDelayedServer(11 * time.Second)
			serverB := makeDelayedServer(12 * time.Second)

			_, err := racer.Racer(serverA.URL, serverB.URL)

			if err == nil {
				t.Error("expected an error but did not get one")
			}
		})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
