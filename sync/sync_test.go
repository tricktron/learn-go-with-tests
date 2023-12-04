package sync_test

import (
	"testing"

	"learn-go-with-tests/sync"
)

func TestCounter(t *testing.T) {
	t.Parallel()
	t.Run("incrementing the counter 3 times leaves it at 3",
		func(t *testing.T) {
			t.Parallel()
			counter := sync.Counter{}
			counter.Inc()
			counter.Inc()
			counter.Inc()

			if counter.Value() != 3 {
				t.Errorf("got %d want %d", counter.Value(), 3)
			}
		},
	)
}
