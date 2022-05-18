package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println(people)

	a := []int{10, 2, 8, 20, 5, 3, 9, 1, 32, 0, 45}
	// sort.Ints(a)
	sort.SliceStable(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Println(a)

}
