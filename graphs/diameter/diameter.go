package diameter

import "container/list"

// pairs |- [][]int represent undirected edge between pair u->v for all pairs[i]
// returns the maximum diameter i.e maximum count of number of edges between any two node
func DiameterInUndirectedGraph(pairs [][]int) int {
	// 1. build a graph from the given pairs
	n := len(pairs) + 1
	graph := make([][]int, n)

	for _, pair := range pairs {
		from, to := pair[0], pair[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	// 1. pick any node, let's say 0
	// 2. apply bfs on the 0 to get the farthest node of one end
	// 3. apply bfs on the node from 2nd step to the get the farthest side

	randomNode := 0
	lastNode, _ := bsf(graph, randomNode)

	_, totalMaximumPath := bsf(graph, lastNode)

	return totalMaximumPath
}

func bsf(graph [][]int, startNode int) (lastNode int, totalEdges int) {
	visited := make([]bool, len(graph))

	q := list.New()
	q.PushBack(startNode)
	visited[startNode] = true

	for q.Len() > 0 {
		qLen := q.Len()
		for range qLen {
			node := q.Remove(q.Front()).(int)
			lastNode = node
			// visit all child of the node
			for _, child := range graph[node] {
				if !visited[child] {
					visited[child] = true
					q.PushBack(child)
				}
			}
		}
		if q.Len() > 0 {
			totalEdges++
		}
	}

	return
}
