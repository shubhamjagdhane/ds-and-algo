package eulerpath

type graphFormat int

const (
	None graphFormat = iota
	EulerPath
	EulerCircuit
)

func IsEulerCircuitInUndirectedGraph(n int, graph [][]int) int {
	// 1. Check whether the graph is connected or not
	if !isConnected(n, graph) {
		return int(None)
	}

	numberOfOddDegreeVertices := 0
	for i := range graph {
		// if degree of vertex is odd then
		// it will not form a Euler-Circuit
		if len(graph[i])&1 == 1 {
			numberOfOddDegreeVertices++
		}
	}

	if numberOfOddDegreeVertices == 1 || numberOfOddDegreeVertices > 2 {
		// if odd degree vertices are one or greater than two
		// graph neigther exists EulerPath nor EulerCircuit
		return int(None)
	}

	if numberOfOddDegreeVertices == 2 {
		// if exactly two vertices have odd degree
		// then graph has a EulerPath
		return int(EulerPath)
	}

	// which means all vertices have a even degree
	// indicates graph has a EulerPath as well as EulerCircuit
	return int(EulerCircuit)
}

func isConnected(n int, graph [][]int) bool {
	// finding non-zero degree vertex
	nonZeroDegreeVertex := -1
	for i := range n {
		if len(graph[i]) != 0 {
			nonZeroDegreeVertex = i
			break
		}
	}

	// basic dfs tarversal on graph
	var dfs func(node int, visited []bool)
	dfs = func(node int, visited []bool) {
		visited[node] = true
		for _, child := range graph[node] {
			if !visited[child] {
				dfs(child, visited)
			}
		}
	}
	visited := make([]bool, n)
	// traversing graph from nonZeroDegreeVertex
	// so that it will reach to all connected vertices
	dfs(nonZeroDegreeVertex, visited)

	// checking whether any nonZeroDegreeVertex not visited
	// if not visited then it is part of different connected component
	// which means graph is not connected
	for i := range n {
		if !visited[i] && len(graph[i]) > 0 {
			return false
		}
	}

	// if no any non-zero degree vertex found that not visited
	// which means all vertices have been visited by dfs call
	return true
}
