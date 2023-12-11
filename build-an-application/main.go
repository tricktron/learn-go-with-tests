package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/build-an-application/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(_ string) int {
	return 123 //nolint: gomnd
}

func (i *InMemoryPlayerStore) RecordWin(_ string) {}

func main() {
	server := &server.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server)) //nolint: gosec
}
