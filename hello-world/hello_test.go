package main

import "testing"

func TestHello(t *testing.T) {
    got  := Hello("Michael")
    want := "Hello, Michael"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
