package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(writer http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	if playerName == "Pepper" {
		fmt.Fprint(writer, "20")

		return
	}

	if playerName == "Floyd" {
		fmt.Fprint(writer, "10")

		return
	}
}
