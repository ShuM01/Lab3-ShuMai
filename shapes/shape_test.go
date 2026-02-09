package main

import "testing"

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	expected := 20.0
	if r.Area() != expected {
		t.Errorf("got %.2f, expected %.2f", r.Area(), expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 3}
	expected := 2 * 3.14159 * 3
	if c.Perimeter() < expected-0.01 || c.Perimeter() > expected+0.01 {
		t.Errorf("got %.2f, expected around %.2f", c.Perimeter(), expected)
	}
}

func TestTriangleScale(t *testing.T) {
	tg := Triangle{Base: 6, Height: 4}
	tg.Scale(2)
	if tg.Base != 12 || tg.Height != 8 {
		t.Errorf("Scale failed, got Base=%.2f Height=%.2f", tg.Base, tg.Height)
	}
}
