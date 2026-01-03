package topologicalsort

import "container/list"

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected directed graph where ith indicates children of ith node
func TopologiclSortByDFS(n int, graph [][]int) []int {
	stack := list.New()

	var dfs func(node int, visited []bool)
	dfs = func(node int, visited []bool) {
		visited[node] = true
		for _, child := range graph[node] {
			if !visited[child] {
				dfs(child, visited)
			}
		}
		stack.PushBack(node)
	}

	visited := make([]bool, n)

	// traversing each node as there might be vertex which is not
	// reachable from other vertex
	for i := range n {
		if !visited[i] {
			dfs(i, visited)
		}
	}

	result := make([]int, 0, stack.Len())
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(int)
		result = append(result, node)
	}

	return result
}

// n |- int represent the total number of nodes in the graph
// graph |- [][]int connected directed graph where ith indicates children of ith node
func TopologiclSortByBFS(n int, graph [][]int) []int {
	inDegree := make([]int, n)
	visited := make([]bool, n)

	for _, edges := range graph {
		for _, edge := range edges {
			inDegree[edge]++
		}
	}
	q := list.New()
	for node, degree := range inDegree {
		if degree == 0 {
			q.PushBack(node)
		}
	}
	result := make([]int, 0, n)
	for q.Len() > 0 {
		node := q.Remove(q.Front()).(int)
		if visited[node] {
			continue
		}
		result = append(result, node)
		visited[node] = true
		for _, child := range graph[node] {
			if !visited[child] {
				q.PushBack(child)
			}
		}
	}

	return result
}
