package go_specs_greet //nolint: revive,stylecheck

import (
	"io"
	"net/http"
)

type Driver struct {
	BaseURL string
}

func (d Driver) Greet() (string, error) {
	res, err := http.Get(d.BaseURL + "/greet") //nolint: noctx
	if err != nil {
		return "", err //nolint: wrapcheck
	}
	defer res.Body.Close()

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err //nolint: wrapcheck
	}

	return string(greeting), nil
}
