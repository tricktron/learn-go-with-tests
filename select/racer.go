package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	url1Duration := measureResponseTime(url1)
	url2Duration := measureResponseTime(url2)

	if url1Duration < url2Duration {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()

	http.Get(url) //nolint: errcheck,gosec,bodyclose,noctx

	return time.Since(start)
}
