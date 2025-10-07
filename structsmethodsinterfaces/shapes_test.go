package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("area of rectangles", func(t *testing.T) {
		rectangle := Rectangle{5, 8}
		checkArea(t, rectangle, 40.00)
	})
	t.Run("area of circles", func(t *testing.T) {
		circle := Circle{4}
		checkArea(t, circle, 50.26548245743669)

	})
}
