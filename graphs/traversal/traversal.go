package traversal

import "container/list"

// n |- int represent the total number of nodes in the graph
// graph |- [][]int represent children of ith node at ith index
func dfs(n int, graph [][]int) []int {
	result := make([]int, 0, n)
	visited := make([]bool, n)

	var solve func(node int, visited []bool)
	solve = func(node int, visited []bool) {
		visited[node] = true
		result = append(result, node)
		for _, child := range graph[node] {
			if !visited[child] {
				solve(child, visited)
			}
		}
	}

	// starting with any node for undirected connected graph
	// if the given graph is not connected then, loop through on each vertex
	solve(0, visited)

	return result
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int represent children of ith node at ith index
func bfs(n int, graph [][]int) []int {
	result := make([]int, 0, n)
	visited := make([]bool, n)

	q := list.New()
	q.PushBack(0) // starting from 0 node

	for q.Len() > 0 {
		node := q.Remove(q.Front()).(int)
		if visited[node] {
			continue
		}
		visited[node] = true
		result = append(result, node)
		for _, child := range graph[node] {
			if !visited[child] {
				q.PushBack(child)
			}
		}
	}

	return result
}
