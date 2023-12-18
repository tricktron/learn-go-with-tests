package interactions_test

import (
	"testing"

	"learn-go-with-tests/go-specs-greet/domain/interactions"
	"learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
