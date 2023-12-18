package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil { //nolint:gosec
		log.Fatal(err)
	}
}
