package go_specs_greet_test

import (
	"testing"

	go_specs_greet "learn-go-with-tests/go-specs-greet"
	"learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_specs_greet.Greet),
	)
}
