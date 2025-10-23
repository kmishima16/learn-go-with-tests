package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("長方形の周囲の長さ", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		got := rect.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("長方形の面積", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		checkArea(t, rect, 100.0)
	})

	t.Run("円の面積", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
