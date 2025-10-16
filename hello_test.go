package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("名前を指定して挨拶できる", func(t *testing.T) {
		got := Greet("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("名前が空の場合、デフォルト値を使う", func(t *testing.T) {
		got := Greet("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
