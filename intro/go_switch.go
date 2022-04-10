package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	daw := rand.Intn(6) + 1
	fmt.Println("Day", daw)

	x := 42
	result := ""
	switch x {
	case 1:
		result = "It is Sunday"
	case 7:
		result = "It is Saturday"
	default:
		result = "It is week day"
	}

	fmt.Println(result)

}
