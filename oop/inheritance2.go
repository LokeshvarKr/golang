package main

import "fmt"


// inheritance (embedding and interface)
type IBase interface {
	SayHello()
}

type Base struct {
	Color string
}

func (b Base) SayHello() {
	fmt.Println("Hello From Base")
	b.SayHi()
}

func (b Base) SayHi(){
	fmt.Println("Hi From Base")
}
type Child struct {
	Base  //embedding
	Style string
}

func (b Child) SayHi(){
	fmt.Println("Hi From Child")
}

func Hello(b IBase) {
	b.SayHello()
}

func main() {

	b := Base{
		Color: "Green",
	}
	c := Child{
		Base:  b,
		Style: "simple",
	}

	Hello(b)
	Hello(c) // <--- from Hello() SayHello() is called and inised SayHello() SayHi() is called and incase of Child 
	// object Child implementation of SayHi() should be called but it won't, Base implementation will be called here

	// One way to fix the above problem is to make SayHi() as a property which is of type function
	// in the base struct. This is possible in GO as functions are first-class variables in Go
}
