package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("5回繰り返す", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
