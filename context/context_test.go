package servercontext_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server_context "learn-go-with-tests/context"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	t.Parallel()

	data := "hello, world"
	server := server_context.Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
