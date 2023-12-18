package go_specs_greet //nolint: revive,stylecheck

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello %s", req.URL.Query().Get("name"))
}
