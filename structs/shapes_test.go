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

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Circle{Radius: 1}, want: 3.141592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run("", func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}
