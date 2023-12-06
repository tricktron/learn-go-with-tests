package servercontext_test

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	server_context "learn-go-with-tests/context"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true

	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true

	return 0, errors.New("not implemented") //nolint: goerr113
}

func (s *SpyResponseWriter) WriteHeader(_ int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, char := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")

				return
			default:
				time.Sleep(10 * time.Millisecond)

				result += string(char)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err() //nolint: wrapcheck
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Parallel()

	t.Run("request without cancelling returns data from store", func(t *testing.T) {
		t.Parallel()
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		server := server_context.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("request with cancelling does not fetch from store", func(t *testing.T) {
		t.Parallel()
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		server := server_context.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{written: false}

		server.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
