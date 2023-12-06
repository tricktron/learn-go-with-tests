package servercontext

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
