package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Parallel()

	t.Run("Countdown prints 3 2 1 Go!", func(t *testing.T) {
		t.Parallel()

		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{Calls: 0}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		t.Parallel()

		spySleeperPrinter := &SpyCountdownOperations{Calls: []string{}}

		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(_ []byte) (int, error) {
	s.Calls = append(s.Calls, write)

	return len(s.Calls), nil
}

const (
	sleep = "sleep"
	write = "write"
)
