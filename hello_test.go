package main

import "testing"



func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("挨拶の言語を指定できる（フランス語）", func(t *testing.T) {
		got := Hello("Fou", "French")
		want := "Bonjour, Fou"

		assertMessage(t, got, want)
	})
	t.Run("挨拶の言語を指定できる（フランス語）", func(t *testing.T) {
		got := Hello("Fou", "French")
		want := "Bonjour, Fou"

		assertMessage(t, got, want)
	})
	t.Run("挨拶の対象（名前）を指定できる", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertMessage(t, got, want)
	})
	t.Run("名前が空の場合、デフォルトの挨拶を返す", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
	t.Run("言語が空の場合、英語を返す", func(t *testing.T) {
		got := Hello("test", "")
		want := "Hello, test"

		assertMessage(t, got, want)
	})
}
