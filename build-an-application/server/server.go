package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		writer.WriteHeader(http.StatusAccepted)
	}

	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(playerName)

	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(writer, score)
}
