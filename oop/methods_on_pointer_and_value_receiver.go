package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

// SetAge is a method on pointer receiver
func (e *Employee) SetAge(age int) {
	e.Age = age
}

// SetName is a method on value receiver
func (e Employee) SetName(name string) {
	e.Name = name
}
func main() {
	e := Employee{
		Name: "Teegen",
		Age:  22,
	}
	fmt.Printf("%#v\n", e)
	e.SetAge(23)
	fmt.Printf("%#v\n", e)

	// becaue SetName is value receiver method it wont be able to change Name in object e
	e.SetName("Teegen Zongna")
	fmt.Printf("%#v\n", e)
}
