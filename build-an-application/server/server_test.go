package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-go-with-tests/build-an-application/server"
)

func TestGetPlayers(t *testing.T) {
	t.Parallel()

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &server.PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		t.Parallel()
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		t.Parallel()
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		t.Parallel()
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	//nolint: noctx
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func assertResponseBody(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}
