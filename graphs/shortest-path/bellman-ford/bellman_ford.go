package bellmanford

import "math"

type Node struct {
	vertex, wt int
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected undirected edges where ith indicates children of ith node and edge[i] = [u, v, w] represents the edge between the nodes u and v having w edge weight.
func shortestPath(n int, edges [][]int, src int) []int {
	// 0:[Node{v, wt}]
	distance := make([]int, n)
	for idx := range distance {
		distance[idx] = math.MaxInt
	}

	// starting from the given source
	distance[src] = 0

	for range n { // relaxation for (n-1) times
		for _, edge := range edges {
			u, v, wt := edge[0], edge[1], edge[2]
			if distance[u] != math.MaxInt &&
				distance[u]+wt < distance[v] {
				distance[v] = distance[u] + wt
			}
		}
	}

	return distance
}
