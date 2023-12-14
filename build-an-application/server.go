package main

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

func (p *PlayerServer) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

	router.Handle("/players/", http.HandlerFunc(
		func(_ http.ResponseWriter, req *http.Request) {
			playerName := strings.TrimPrefix(req.URL.Path, "/players/")

			switch req.Method {
			case http.MethodPost:
				p.processWin(writer, playerName)
			case http.MethodGet:
				p.showScore(writer, playerName)
			}
		}))
	router.ServeHTTP(writer, req)
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
