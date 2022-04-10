package main

import "fmt"

func main() {
	x := 2.0
	y := 4.0

	if f := x / y; f > 0.5 {
		fmt.Printf("f : %v is greater than 0.5\n", f)
	} else {
		fmt.Printf("f : %v is less than or equal to 0.5\n", f)
	}
}
