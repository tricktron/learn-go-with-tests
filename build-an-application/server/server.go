package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(writer http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(writer, getPlayerScore(playerName))
}

func getPlayerScore(name string) string {
	playerScore := ""
	if name == "Pepper" {
		playerScore = "20"
	}

	if name == "Floyd" {
		playerScore = "10"
	}

	return playerScore
}
