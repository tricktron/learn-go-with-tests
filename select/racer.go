package racer

import (
	"net/http"
)

func Racer(url1, url2 string) string {
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		http.Get(url) //nolint: errcheck,gosec,bodyclose,noctx
		close(channel)
	}()

	return channel
}
