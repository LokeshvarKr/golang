package main

import "fmt"

func main() {
	var x int
	var y float64

	x, y = 10, 20.0

	fmt.Printf("x : %v Type : %T\n", x, x)
	fmt.Printf("y : %v Type : %T\n", y, y)

	var mean float64

	mean = (float64(x) + y) / 2.0

	fmt.Printf("mean : %v and type : %T\n", mean, mean)

}
