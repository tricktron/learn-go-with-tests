package racer

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

var errTimeout = errors.New("timed out waiting")

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
		return "", makeTimeoutErr(url1, url2)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		res, _ := http.Get(url) //nolint: gosec,noctx
		res.Body.Close()
		close(channel)
	}()

	return channel
}

func makeTimeoutErr(url1, url2 string) error {
	return fmt.Errorf("%w for %s and %s", errTimeout, url1, url2)
}
