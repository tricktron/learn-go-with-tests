package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-go-with-tests/build-an-application/server"
)

func TestGetPlayers(t *testing.T) {
	t.Parallel()
	t.Run("returns Pepper's score", func(t *testing.T) {
		t.Parallel()
		//nolint: noctx
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
