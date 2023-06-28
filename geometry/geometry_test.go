package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Perimeter()

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, &rectangle, 40.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, &circle, 2*math.Pi*circle.Radius)
	})

}

func TestArea(t *testing.T) {

	t.Helper()

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: &Rectangle{10.0, 10.0}, want: 100.0},
		{shape: &Circle{10.0}, want: math.Pi * 10.0 * 10.0},
		{shape: &Triangle{10.0, 10.0}, want: 50.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("want %.2f, got %.2f, received %#v", tt.want, got, tt.shape)
		}
	}

}
