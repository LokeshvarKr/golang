package main

import "fmt"

func printGraph(graph [][]int, vertices int) {
	for v := 0; v < vertices; v++ {
		fmt.Printf("%d: ", v)
		for e := 0; e < len(graph[v]); e++ {
			fmt.Printf("%d ", graph[v][e])
		}
		fmt.Printf("\n")
	}
}
func main() {
	fmt.Println("Hello!")
	var vertices int
	fmt.Scanf("%d\n", &vertices)
	graph := make([][]int, vertices)
	for v := 0; v < vertices; v++ {
		graph[v] = make([]int, 0)
	}
	var edges int
	fmt.Scanf("%d\n", &edges)

	for e := 0; e < edges; e++ {
		var i int
		var j int
		fmt.Scanf("%d %d\n", &i, &j)
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}
	printGraph(graph, vertices)
}
