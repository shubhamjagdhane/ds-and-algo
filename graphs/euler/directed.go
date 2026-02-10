package eulerpath

func IsEulerCircuitInDirectedGraph(n int, graph [][]int) int {
	// check whether the given graph is connected or not
	if !isConnected(n, graph) {
		return int(None)
	}

	indegree, outdegree := make([]int, n), make([]int, n)

	// calculating indegree and outdegree from a given graph
	for node := range graph {
		outdegree[node] = len(graph[node])
		for _, child := range graph[node] {
			indegree[child]++
		}
	}

	var startNodeCount, endNodeCount, remainingNodes int
	for idx := range n {
		if outdegree[idx]-indegree[idx] == 1 {
			// start node
			startNodeCount++
		}
		if indegree[idx]-outdegree[idx] == 1 {
			endNodeCount++
		}

		if indegree[idx] == outdegree[idx] {
			remainingNodes++
		}
	}

	// if every node has the equal indegree & outdegree
	// then given has a the Euler Circuit
	if remainingNodes == n {
		return int(EulerCircuit)
	}

	// if there are multiple startNode or endNode
	// then the graph neigther has Euler Circuit or Euler Path
	if startNodeCount > 1 || endNodeCount > 1 {
		return int(None)
	}

	// if sum of remainingNodes, startNodeCount & endNodeCount
	// not equal to given number of nodes
	// then the graph neigther has Euler Circuit or Euler Path
	if (remainingNodes + startNodeCount + endNodeCount) != n {
		return int(None)
	}

	return int(EulerPath)
}
