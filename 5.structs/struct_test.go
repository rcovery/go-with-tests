package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeter := Perimeter(10.0, 10.0)
	want := 40.0

	if perimeter != want {
		t.Errorf("got %.2f want %.2f", perimeter, want)
	}
}

func TestArea(t *testing.T) {
	area := Area(5, 5)
	want := 25.0

	if area != want {
		t.Errorf("got %.2f want %.2f", area, want)
	}
}
