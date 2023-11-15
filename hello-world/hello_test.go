package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("Hello should return greeting with name", func(t *testing.T) {
        got  := Hello("Michael")
        want := "Hello, Michael"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("Hello without name should return 'Hello World'", func(t *testing.T) {
        got   := Hello("")
        want  := "Hello, World"
        
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
}
