package syncer_test

import (
	"sync"
	"testing"

	syncer "learn-go-with-tests/sync"
)

func TestCounter(t *testing.T) {
	t.Parallel()
	t.Run("incrementing the counter 3 times leaves it at 3",
		func(t *testing.T) {
			t.Parallel()
			counter := syncer.NewCounter()
			counter.Inc()
			counter.Inc()
			counter.Inc()

			assertCounter(t, counter, 3)
		},
	)

	t.Run("runs safely concurrently", func(t *testing.T) {
		t.Parallel()
		wantedCount := 1000
		counter := syncer.NewCounter()

		var waitGroup sync.WaitGroup
		waitGroup.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(tb testing.TB, counter *syncer.Counter, want int) {
	tb.Helper()

	counterValue := counter.Value()
	if counterValue != want {
		tb.Errorf("got %d want %d", counterValue, want)
	}
}
