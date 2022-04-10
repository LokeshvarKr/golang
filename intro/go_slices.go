package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	num := append(numbers, 4)

	fmt.Println(numbers)
	fmt.Println(num)

	// fmt.Println(&numbers)
	// fmt.Println(&num)

	num2 := append(numbers[1:])
	num3 := append(numbers[:len(numbers)-1])

	fmt.Println(num2)
	fmt.Println(num3)

	fmt.Println("Creating slices whit make()")
	a := make([]int, 0)
	fmt.Println(a, len(a), cap(a))
	b := make([]int, 0, 5)
	fmt.Println(b, len(b), cap(b))
	c := make([]int, 5, 5)
	fmt.Println(c, len(c), cap(c))

	a = append(a, 5)
	fmt.Println(a, len(a), cap(a))

	d := append(c, -1)
	fmt.Println(d, len(d), cap(d))
	sort.Ints(d)
	fmt.Println(d, len(d), cap(d))

}
