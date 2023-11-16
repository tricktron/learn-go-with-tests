package main

import "testing"

func TestHello(t *testing.T) {
	t.Parallel()
	t.Run("Hello should return greeting with name", func(t *testing.T) {
		t.Parallel()
		got := Hello("Michael", "English")
		want := "Hello, Michael"
		assertHelloMessage(t, got, want)
	})

	t.Run("Hello without name should return 'Hello World'", func(t *testing.T) {
		t.Parallel()
		got := Hello("", "English")
		want := "Hello, World"
		assertHelloMessage(t, got, want)
	})

	t.Run("Hello in Spanish should return 'Hola' greeting", func(t *testing.T) {
		t.Parallel()
		got := Hello("Pamela", "Spanish")
		want := "Hola, Pamela"
		assertHelloMessage(t, got, want)
	})

	t.Run("Hello in French should return 'Bonjour' greeting", func(t *testing.T) {
		t.Parallel()
		got := Hello("Marguerite", "French")
		want := "Bonjour, Marguerite"
		assertHelloMessage(t, got, want)
	})

	t.Run("Hello in German should return 'Hallo' greeting", func(t *testing.T) {
		t.Parallel()
		got := Hello("David", "German")
		want := "Hallo, David"
		assertHelloMessage(t, got, want)
	})
}

func assertHelloMessage(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q want %q", got, want)
	}
}
