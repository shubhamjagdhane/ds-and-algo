package dijkstra

import (
	"container/heap"
	"math"
)

type Node struct {
	vertex, wt int
}
type MinHeap []Node

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i].wt < mh[j].wt }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(val any)      { *mh = append(*mh, val.(Node)) }
func (mh *MinHeap) Pop() any {
	old := *mh
	n := len(old)
	temp := old[n-1]
	*mh = old[:n-1]
	return temp
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected undirected edges where ith indicates children of ith node and edge[i] = [u, v, w] represents the edge between the nodes u and v having w edge weight.
// source |- target is to find the shortest distance of all the vertices from the source vertex src
func shortestPath(n int, edges [][]int, src int) []int {

	// converting to undirected graph
	graph := make([][]Node, n)
	for _, edge := range edges {
		from, to, wt := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], Node{vertex: to, wt: wt})
	}

	visited := make([]bool, n)

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	heap.Push(minHeap, Node{vertex: src, wt: 0})

	// to find minimum distance from source
	// considering each node it at the maximum distance
	result := make([]int, n)
	for idx := range result {
		result[idx] = math.MaxInt
	}

	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		if visited[node.vertex] {
			continue
		}
		visited[node.vertex] = true
		if result[node.vertex] > node.wt {
			result[node.vertex] = node.wt
		}
		for _, child := range graph[node.vertex] {
			if !visited[child.vertex] {
				heap.Push(minHeap, Node{vertex: child.vertex, wt: node.wt + child.wt})
			}
		}
	}

	return result
}
