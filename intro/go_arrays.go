package main

import (
	"fmt"
)

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[3])
	fmt.Println(len(numbers))
	fmt.Println(numbers)
	fmt.Println(numbers[1:3])

	var colors [3]string
	fmt.Println(colors)
	colors[0] = "Blue"
	colors[2] = "Red"
	fmt.Println(colors)
	fmt.Println(len(colors))

}
