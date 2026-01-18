package flyodwarshall

import "math"

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected undirected graph represents the weight of the edge from node i to node j.
func ShortestPath(n int, graph [][]int) [][]int {
	rows, cols := len(graph), len(graph[0])

	for via := range n {
		for i := range rows {
			for j := range cols {
				if !(graph[i][via] == math.MaxInt || graph[via][j] == math.MaxInt) {
					graph[i][j] = min(graph[i][j], graph[i][via]+graph[via][j])
				}
			}
		}
	}

	return graph
}
