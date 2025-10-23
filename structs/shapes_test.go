package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("長方形の周囲の長さ", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("長方形の面積", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
