package main

import "fmt"

// inheritance (embedding and interface)
type IBase interface {
	SayHello()
}

type Base struct {
	Color string
}
type Child struct {
	Base  //embedding
	Style string
}

func (b Base) SayHello() {
	fmt.Println("Hello From Base")
}

func Hello(b IBase) {
	b.SayHello()
}

func main() {

	// c := Child{
	// 	Style: "simple",
	// }
	// c.PrintColor()

	b := Base{
		Color: "Green",
	}
	// c := Child{
	// 	Base:  b,
	// 	Style: "simple",
	// }

	// Hello(b)
	// Hello(c)
	c := Child{
		Base: b,
		Style: "simple",
	}
	Hello(c)
}
