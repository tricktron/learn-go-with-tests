package main

import (
	"log"
	"net/http"

	go_specs_greet "learn-go-with-tests/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil { //nolint:gosec
		log.Fatal(err)
	}
}
