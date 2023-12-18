package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "learn-go-with-tests/go-specs-greet"
)

func Handler(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprint(writer, go_specs_greet.Greet(name))
}
