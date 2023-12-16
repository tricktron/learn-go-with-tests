package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet() (string, error)
}

func GreetSpecification(tb testing.TB, greeter Greeter) {
	tb.Helper()

	got, err := greeter.Greet()
	assert.NoError(tb, err)
	assert.Equal(tb, got, "Hello World")
}
