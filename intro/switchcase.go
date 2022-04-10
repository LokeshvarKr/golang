package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("Enter value of x :\n", &x)
	switch x {
	case 1:
		fmt.Printf("x : %v", x)
	case 2:
		fmt.Printf("x : %v", x)
	default:
		fmt.Printf("x : %v", x)
	}
}
