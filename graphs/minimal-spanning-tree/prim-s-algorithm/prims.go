package primsalgorithm

import (
	"container/heap"
)

type Node struct {
	val, wt int
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
func minimumSpanningTree(n int, edges [][]int) int {
	graph := make([][]Node, n)
	for _, edge := range edges {
		from, to, wt := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], Node{val: to, wt: wt})
		graph[to] = append(graph[to], Node{val: from, wt: wt})
	}
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// starting from the node 0
	heap.Push(minHeap, Node{val: 0, wt: 0})

	minDist := 0
	visited := make([]bool, n)
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		if visited[node.val] {
			continue
		}

		visited[node.val] = true
		minDist += node.wt
		for _, child := range graph[node.val] {
			if !visited[child.val] {
				heap.Push(minHeap, child)
			}
		}
	}

	return minDist
}
