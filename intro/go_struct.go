package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poodle := Dog{"Poodle", 24}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v,Weight: %v", poodle.Breed, poodle.Weight)
}
