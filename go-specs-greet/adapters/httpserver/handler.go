package httpserver

import (
	"fmt"
	"net/http"

	"learn-go-with-tests/go-specs-greet/domain/interactions"
)

func Handler(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Greet(name))
}
