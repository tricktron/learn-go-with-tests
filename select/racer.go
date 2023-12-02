package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	startURL1 := time.Now()
	http.Get(url1)
	url1Duration := time.Since(startURL1)

	startURL2 := time.Now()
	http.Get(url2)
	url2Duration := time.Since(startURL2)

	if url1Duration < url2Duration {
		return url1
	}

	return url2
}
