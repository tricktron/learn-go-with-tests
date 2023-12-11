package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
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
