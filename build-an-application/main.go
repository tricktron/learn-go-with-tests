package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/build-an-application/server"
)

func main() {
	server := &server.PlayerServer{Store: nil}
	log.Fatal(http.ListenAndServe(":5000", server)) //nolint: gosec
}
