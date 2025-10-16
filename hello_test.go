package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("あいさつを返す", func(t *testing.T) {
		got := Hello("Ken", "English")
		want := "Hello, Ken"

		assertMessage(t, got, want)
	})

	t.Run("名前が空文字の場合はWorldにあいさつを返す", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertMessage(t, got, want)
	})

	t.Run("言語が空文字の場合は英語であいさつを返す", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"

		assertMessage(t, got, want)
	})

	t.Run("挨拶の言語を指定できる（日本語）", func(t *testing.T) {
		got := Hello("Taro", "Japanese")
		want := "こんにちは、Taro"

		assertMessage(t, got, want)
	})

	t.Run("挨拶の言語を指定できる（スペイン語）", func(t *testing.T) {
		got := Hello("Carlos", "Spanish")
		want := "Hola, Carlos"

		assertMessage(t, got, want)
	})
}
