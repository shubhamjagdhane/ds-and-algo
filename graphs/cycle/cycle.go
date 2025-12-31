package cycle

import "container/list"

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected undirected graph where ith indicates children of ith node
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
// graph |- [][]int connected undirected graph where ith indicates children of ith node
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

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected directed graph where ith indicates children of ith node
func isCycleDFSInDirectedGraph(n int, graph [][]int) bool {
	visited := make([]bool, n)
	inRecursion := make([]bool, n)

	var dfs func(node int, visited, inRecursion []bool) bool
	dfs = func(node int, visited, inRecursion []bool) bool {
		visited[node] = true
		inRecursion[node] = true
		for _, child := range graph[node] {
			if !visited[child] && dfs(child, visited, inRecursion) {
				return true
			} else if inRecursion[child] {
				return true
			}
		}
		inRecursion[node] = false
		return false
	}

	return dfs(0, visited, inRecursion)
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected directed graph where ith indicates children of ith node
func isCycleBFSInDirectedGraph(n int, graph [][]int) bool {
	visited := make([]bool, n)
	inDegree := make([]int, n)
	for _, edges := range graph {
		for _, edge := range edges {
			inDegree[edge]++
		}
	}

	q := list.New()
	for node, val := range inDegree {
		if val == 0 {
			q.PushBack(node)
		}
	}
	result := make([]int, 0, n)
	for q.Len() > 0 {
		node := q.Remove(q.Front()).(int)
		if visited[node] {
			return false
		}
		visited[node] = true
		result = append(result, node)
		for _, child := range graph[node] {
			if visited[child] {
				return true
			}
			q.PushBack(child)
		}
	}

	if len(result) == n {
		return false
	}

	return true
}
