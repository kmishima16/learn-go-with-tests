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
		got := Hello("World", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})

	t.Run("挨拶の相手を指定できる", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertMessage(t, got, want)
	})

	t.Run("空文字の場合はデフォルトのあいさつを返す", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})

	t.Run("挨拶の言語を指定できる（日本語）", func(t *testing.T) {
		got := Hello("Taro", "Japanese")
		want := "こんにちは、Taro"

		assertMessage(t, got, want)
	})
}
