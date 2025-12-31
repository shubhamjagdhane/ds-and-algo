package cycle

import "container/list"

// n |- int represent the total number of nodes in the graph
// graph |- [][]int represent children of ith node at ith index
func isCycleDFSInUndirectedGraph(n int, graph [][]int) bool {
	visited := make([]bool, n)

	var dfs func(node int, parent int, visited []bool) bool
	dfs = func(node int, parent int, visited []bool) bool {
		visited[node] = true
		for _, child := range graph[node] {
			if child == parent {
				continue
			}
			if visited[child] {
				return true
			}
			if dfs(child, node, visited) {
				return true
			}
		}

		return false
	}

	return dfs(0, -1, visited)
}

type Node struct {
	vertex, parent int
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int represent children of ith node at ith index
func isCycleBFSInUndirectedGraph(n int, graph [][]int) bool {
	q := list.New()
	q.PushBack(Node{0, -1})
	visited := make([]bool, n)

	for q.Len() > 0 {
		node := q.Remove(q.Front()).(Node)
		if visited[node.vertex] {
			return true
		}
		visited[node.vertex] = true
		for _, child := range graph[node.vertex] {
			if child == node.parent {
				continue
			}
			if visited[child] {
				return true
			}
			q.PushBack(Node{vertex: child, parent: node.vertex})
		}
	}

	return false
}
