package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf( //nolint: goerr113
			"timed out waiting for %s and %s",
			url1,
			url2,
		)
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
