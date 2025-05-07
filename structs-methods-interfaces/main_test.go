package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Cricle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 6.0},
	}

	for _, test := range areaTest {
		area := test.shape.Area()

		t.Run(test.name, func(t *testing.T) {
			if test.hasArea != area {
				t.Errorf("%#v want %g but got %g", test.shape, test.hasArea, area)
			}
		})

	}
}
