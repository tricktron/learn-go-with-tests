package app

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "20")
}
