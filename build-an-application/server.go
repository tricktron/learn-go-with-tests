package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	playerServer := new(PlayerServer)

	playerServer.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(playerServer.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(playerServer.playersHandler))

	playerServer.Handler = router

	return playerServer
}

func (p *PlayerServer) leagueHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("content-type", jsonContentType)
	json.NewEncoder(writer).Encode(p.Store.GetLeague()) //nolint: errcheck,errchkjson
}

func (p *PlayerServer) playersHandler(writer http.ResponseWriter, req *http.Request) {
	playerName := strings.TrimPrefix(req.URL.Path, "/players/")

	switch req.Method {
	case http.MethodPost:
		p.processWin(writer, playerName)
	case http.MethodGet:
		p.showScore(writer, playerName)
	}
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, playerName string) {
	score := p.Store.GetPlayerScore(playerName)

	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(writer, score)
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, playerName string) {
	p.Store.RecordWin(playerName)
	writer.WriteHeader(http.StatusAccepted)
}
