package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("あいさつを返す", func(t *testing.T) {
		got := Hello("World")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("挨拶の相手を指定できる", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
