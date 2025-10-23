package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("長方形の周囲の長さ", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		got := Perimeter(rect)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("長方形の面積", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		got := Area(rect)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
