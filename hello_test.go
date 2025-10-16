package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("挨拶の言語を指定できる(フランス語)", func(t *testing.T) {
		got := Greet("Chris", "French")
		want := "Bonjour, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("名前を指定して挨拶できる", func(t *testing.T) {
		got := Greet("Chris", "English")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("名前が空の場合、デフォルト値を使う", func(t *testing.T) {
		got := Greet("", "English")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
