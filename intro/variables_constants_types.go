package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)

	var str2 string = "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str2))
	fmt.Println(strings.ToLower(str2))
	fmt.Println(strings.Title(str2))

	const x int = 10
	fmt.Println("const x : ", x)

}
