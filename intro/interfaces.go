package main

import "fmt"

const PI float64 = 3.14

// interface is just an abstract type , a protocol
type Shape interface {
	Area() float64
	Circumference() float64
}

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Circumference() float64 {
	return s.Side * 4.0
}
func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}
func (c Circle) Circumference() float64 {
	return 2.0 * PI * c.Radius
}

func main() {

	shapes := []Shape{
		Square{Side: 2.0},
		Square{Side: 2.5},
		Circle{Radius: 2.0},
		Circle{Radius: 2.5},
	}

	fmt.Printf("type : %T\n",shapes)
	for _, s := range shapes {
		fmt.Printf("type : %T and area : %f\n", s, s.Area())

	}
}
