package kruskalsalgorithm

import (
	"sort"
)

type Node struct {
	val, wt int
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected undirected edges where ith indicates children of ith node and edge[i] = [u, v, w] represents the edge between the nodes u and v having w edge weight.
func minimumSpanningTree(n int, edges [][]int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	visited := make([]bool, n)
	sum := 0
	for _, edge := range edges {
		from, _, wt := edge[0], edge[1], edge[2]
		if visited[from] {
			continue
		}
		visited[from] = true
		sum += wt
	}

	return sum
}
