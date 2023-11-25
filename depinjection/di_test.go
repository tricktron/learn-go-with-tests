package depinjection_test

import (
	"bytes"
	"testing"

	"learn-go-with-tests/depinjection"
)

func TestGreet(t *testing.T) {
	t.Parallel()

	buffer := bytes.Buffer{}
	depinjection.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
