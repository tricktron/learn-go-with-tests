package go_specs_greet //nolint: revive,stylecheck

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprint(writer, Greet(name))
}
