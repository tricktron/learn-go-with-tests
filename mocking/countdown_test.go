package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Parallel()

	t.Run("Countdown prints 3 2 1 Go!", func(t *testing.T) {
		t.Parallel()

		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{Calls: []string{}})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
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

func TestConfigurableSleeper(t *testing.T) {
	t.Parallel()

	sleepTime := 5 * time.Second
	spyTime := &SpyTime{durationSlept: 0}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf(
			"should have slept for %v but slept for %v",
			sleepTime,
			spyTime,
		)
	}
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
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
