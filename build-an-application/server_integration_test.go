package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrieveThem(t *testing.T) {
	t.Parallel()

	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	playerName := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(playerName))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
