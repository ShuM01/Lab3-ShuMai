package main

import "fmt"

func main() {
	r := Rectangle{Width: 4, Height: 5}
	c := Circle{Radius: 3}
	t := Triangle{Base: 6, Height: 4}

	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
	fmt.Printf("Triangle Area: %.2f, Perimeter: %.2f\n", t.Area(), t.Perimeter())

	// Scale example
	r.Scale(2)
	fmt.Printf("Scaled Rectangle: Width=%.2f, Height=%.2f\n", r.Width, r.Height)
}
