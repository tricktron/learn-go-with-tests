package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Parallel()

	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: []string{},
		league:   []Player{},
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		t.Parallel()
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		t.Parallel()
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		t.Parallel()
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	t.Parallel()

	store := StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
		league:   []Player{},
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		t.Parallel()
		playerName := "Pepper"
		request := newPostWinRequest(playerName)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != playerName {
			t.Errorf(
				"did not store correct winner, got %q want %q",
				store.winCalls[0],
				playerName,
			)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Parallel()

	t.Run("it returns 200 on /league", func(t *testing.T) {
		t.Parallel()

		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{
			scores:   map[string]int{},
			winCalls: []string{},
			league:   wantedLeague,
		}
		server := NewPlayerServer(&store)
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response.Body)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertLeague(t, got, wantedLeague)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func newPostWinRequest(name string) *http.Request { //nolint: unparam
	//nolint: noctx
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func newLeagueRequest() *http.Request {
	//nolint: noctx
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)

	return req
}

func getLeagueFromResponse(tb testing.TB, body io.Reader) (league []Player) { //nolint: nonamedreturns
	tb.Helper()

	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		tb.Fatalf(
			"Unable to parse response from server %q into slice of Player '%v'",
			body,
			err,
		)
	}

	return
}

func assertResponseBody(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatusCode(tb testing.TB, got, want int) {
	tb.Helper()

	if got != want {
		tb.Errorf("did not get correct status, got %d want %d", got, want)
	}
}

func assertContentType(
	tb testing.TB,
	response *httptest.ResponseRecorder,
	want string,
) {
	tb.Helper()

	if response.Result().Header.Get("content-type") != want { //nolint: bodyclose
		//nolint: bodyclose
		tb.Errorf(
			"response did not have content-type of %s, got %v",
			want,
			response.Result().Header,
		)
	}
}

func assertLeague(tb testing.TB, got, want []Player) {
	tb.Helper()

	if !reflect.DeepEqual(got, want) {
		tb.Errorf("got %v want %v", got, want)
	}
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}
