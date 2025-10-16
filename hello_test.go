package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("挨拶の対象（名前）を指定できる", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("名前が空の場合、デフォルトの挨拶を返す", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("says hello to the world", func(t *testing.T) {
		got := Hello("World")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
