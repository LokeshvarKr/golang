package main

import "fmt"

// inheritance (embedding and interface)
type IBase interface {
	SayHello()
}

type Base struct {
	Color string
	SayHi func()
}

func (b Base) SayHello() {
	fmt.Println("Hello From Base")
	b.SayHi()
}

type Child struct {
	Base  //embedding
	Style string
}

func Hello(b IBase) {
	b.SayHello()
}

func main() {

	// b := Base{
	// 	Color: "Green",
	// 	SayHi: func() {
	// 		fmt.Println("Hi form Child")
	// 	},
	// }

	c := Child{
		Base: Base{
			Color: "Green",
			SayHi: func() {
				fmt.Println("Hi form Child")
			},
		},
		Style: "simple",
	}
	Hello(c)
}
