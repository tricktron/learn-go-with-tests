package servercontext_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	server_context "learn-go-with-tests/context"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)

	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Parallel()

	t.Run("request without cancelling returns data from store", func(t *testing.T) {
		t.Parallel()
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false, t: t}
		server := server_context.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("request with cancelling does not return data from store", func(t *testing.T) {
		t.Parallel()
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false, t: t}
		server := server_context.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()

	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}
