package main

import "fmt"


// inheritance (embedding)


// Go prefers composition over inheritance.
// It allows embedding of struct into other struct.
// Go does not support type inheritance.

type Base struct {
	Color string
}
type Child struct {
	Base //embedding
	Style string
}

func (b Base) PrintColor() {
	fmt.Printf("Color: %v\n", b.Color)
}

func Hello(b Base){
	fmt.Printf("Color : %s\n",b.Color)
}

func main() {
	
	// c := Child{
	// 	Style: "simple",
	// }
	// c.PrintColor()

	b := Base{
		Color: "Green",
	}
	c := Child{
		Base: b,
		Style: "simple",
	}
	c.PrintColor()

	Hello(b)

	// c := Child{
	// 	Base: b,
	// 	Style: "simple",
	// }
	// Hello(c) // <---- this will give compilation error  as Go does not support type inheritance.
}
