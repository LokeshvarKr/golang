package main

import "fmt"

func printGraph(graph [][]int, vertices int) {
	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			fmt.Printf("%d ", graph[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	var vertices int
	fmt.Scanf("%d\n", &vertices)
	graph := make([][]int, vertices)
	for v := 0; v < vertices; v++ {
		graph[v] = make([]int, vertices)
	}

	var edges int
	fmt.Scanf("%d\n", &edges)
	for e := 0; e < edges; e++ {
		var i int
		var j int
		fmt.Scanf("%d%d\n", &i, &j)
		graph[i][j] = 1
		graph[j][i] = 1
	}
	printGraph(graph, vertices)
}
