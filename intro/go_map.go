package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	states := make(map[string]string)
	states["DL"] = "Delhi"
	states["GA"] = "Goa"
	states["JH"] = "Jharkhan"
	fmt.Println(states)

	delete(states, "JH")
	fmt.Println(states)

	states["AS"] = "Assam"
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//sorting
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted : ")
	for i := range keys {
		fmt.Println(keys[i], states[keys[i]])
	}

	fmt.Printf("%T", states)
}
