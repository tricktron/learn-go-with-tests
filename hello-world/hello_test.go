package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("Hello should return greeting with name", func(t *testing.T) {
        got  := Hello("Michael", "English")
        want := "Hello, Michael"
        assertHelloMessage(t, got, want)
    })

    t.Run("Hello without name should return 'Hello World'", func(t *testing.T) {
        got   := Hello("", "English")
        want  := "Hello, World"
        assertHelloMessage(t, got, want)
    })

    t.Run("Hello in Spanish should return 'Hola' greeting", func(t *testing.T) {
        got   := Hello("Pamela", "Spanish")
        want  := "Hola, Pamela"
        assertHelloMessage(t, got, want)
    })

    t.Run("Hello in French should return 'Bonjour' greeting", func(t *testing.T) {
        got   := Hello("Marguerite", "French")
        want  := "Bonjour, Marguerite"
        assertHelloMessage(t, got, want)
    })
}
        
func assertHelloMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
