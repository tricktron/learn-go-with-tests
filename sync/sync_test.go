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

			assertCounter(t, counter, 3)
		},
	)
}

func assertCounter(tb testing.TB, counter sync.Counter, want int) {
	tb.Helper()

	counterValue := counter.Value()
	if counterValue != want {
		tb.Errorf("got %d want %d", counterValue, want)
	}
}
