package main

import "fmt"

// In compile-time polymorphism, the call is resolved during compile time by the compiler.
// Some of the  forms for compile-time polymorphism are
// (1) Method/Function Overloading: more than one method/function exists with the same name but with
// different signatures or possibly different return types
// (2) Operator Overloading: the Same operator is used for operating on different data types

// Go doesn’t directly support method/function/operator overloading
// but variadic function provides a way of achieving the same with increased code complexity.

type maths struct{}


// add is Variadic function
func (m *maths) add(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func main() {
	m := &maths{}

	fmt.Printf("Result: %d\n", m.add(2, 3))
	fmt.Printf("Result: %d\n", m.add(2, 3, 4))
}