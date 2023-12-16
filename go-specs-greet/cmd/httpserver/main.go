package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, "Hello world")
	})
	if err := http.ListenAndServe(":8080", handler); err != nil { //nolint:gosec
		log.Fatal(err)
	}
}
