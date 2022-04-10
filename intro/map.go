package main

import "fmt"

func main() {
	m := map[string]string{"x": "xx", "y": "yy"}
	fmt.Println(m)
	v, k := m["x"]
	fmt.Println(v, k)
	fmt.Println(m["z"])
}
