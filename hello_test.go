package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to the world", func(t *testing.T) {
		got := Hello()
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
