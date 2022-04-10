package main

import (
	"fmt"
)

func main() {
	var p *int
	if p != nil {
		fmt.Println("value of p: ", *p)
	} else {
		fmt.Println("Value of p is nil")
	}

	var value = 10
	p = &value

	if p != nil {
		fmt.Println("value of p: ", *p)
	} else {
		fmt.Println("Value of p is nil")
	}

	var f float64 = 42.2
	var pp *float64 = &f

	fmt.Println("value of pp is: ", *pp)

}
