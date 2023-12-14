package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrieveThem(t *testing.T) {
	t.Parallel()

	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	playerName := "Pepper"

	t.Run("get score", func(t *testing.T) {
		t.Parallel()

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(playerName))

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		t.Parallel()

		want := []Player{
			{"Pepper", 3},
		}
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newLeagueRequest())
		got := getLeagueFromResponse(t, response.Body)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertLeague(t, got, want)
	})
}
