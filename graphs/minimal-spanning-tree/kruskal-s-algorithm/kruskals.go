package kruskalsalgorithm

import (
	disjointsetunion "ds-and-algo/graphs/disjoint-set-union"
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

	parent, rank := make([]int, n), make([]int, n)

	for idx := range parent {
		parent[idx] = idx
	}
	minDist := 0
	for _, edge := range edges {
		from, to, wt := edge[0], edge[1], edge[2]
		if disjointsetunion.Find(from, parent) != disjointsetunion.Find(to, parent) {
			disjointsetunion.Union(from, to, parent, rank)
			minDist += wt
		}
	}
	return minDist
}
