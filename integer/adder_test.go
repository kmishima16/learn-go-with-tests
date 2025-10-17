package integer

import "testing"

func TestAdd(t *testing.T) {
	t.Run("2つの整数の和を返す", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
