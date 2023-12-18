package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(tb testing.TB, greeter Greeter) {
	tb.Helper()

	got, err := greeter.Greet("Mike")
	assert.NoError(tb, err)
	assert.Equal(tb, got, "Hello Mike")
}
