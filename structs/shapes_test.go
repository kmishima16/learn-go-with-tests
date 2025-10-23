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
		name  string
		shape Shape
		hasArea float64
	}{
		{name: "長方形", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "大きい円", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "小さい円", shape: Circle{Radius: 1}, hasArea: 3.141592653589793},
		{name: "三角形", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %.2f want %.2f", got, tt.hasArea)
			}
		})
	}
}
