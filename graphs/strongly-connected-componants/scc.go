package stronglyconnectedcomponants

import "container/list"

// graph |- [][]int is an adjacency list for a directed graph.
// Index `i` represents node `i`, and `graph[i]` contains a slice of its outgoing neighbors
func KosarajusAlgorithm(graph [][]int) int {
	n := len(graph)
	// 1. topologcalSortUsingDFS
	stack := topologcalSortUsingDFS(n, graph)

	// 2. reverse the given graph i.e. if u->v then make it v->u
	reverseGraph := make([][]int, n)
	for u := range graph {
		for _, v := range graph[u] {
			reverseGraph[v] = append(reverseGraph[v], u)
		}
	}

	// 3. traverse graph in DFS
	var dfs func(node int, visited []bool)
	dfs = func(node int, visited []bool) {
		visited[node] = true
		for _, child := range reverseGraph[node] {
			if !visited[child] {
				dfs(child, visited)
			}
		}
	}
	connectedGraph := 0
	visited := make([]bool, n)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(int)
		if visited[node] {
			continue
		}
		dfs(node, visited)
		connectedGraph++
	}

	return connectedGraph
}

func topologcalSortUsingDFS(n int, graph [][]int) *list.List {
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
	// traverse each node only if they are not already visited
	for i := range n {
		if !visited[i] {
			dfs(i, visited)
		}
	}
	return stack
}

// 1. Travese the graph in Topological way with Stack
// 2. Reverse the given graph
// 3. Pop element from the stack and apply dfs
