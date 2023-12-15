package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrieveLeague(t *testing.T) {
	t.Parallel()

	want := []Player{
		{"Pepper", 3},
		{"Salt", 2},
	}
	pepper := "Pepper"
	salt := "Salt"
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(pepper))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(salt))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(pepper))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(salt))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(pepper))

	getScoreResponse := httptest.NewRecorder()
	server.ServeHTTP(getScoreResponse, newGetScoreRequest(pepper))
	assertStatusCode(t, getScoreResponse.Code, http.StatusOK)
	assertResponseBody(t, getScoreResponse.Body.String(), "3")

	leagueResponse := httptest.NewRecorder()
	server.ServeHTTP(leagueResponse, newLeagueRequest())
	got := getLeagueFromResponse(t, leagueResponse.Body)

	assertLeague(t, got, want)
}
