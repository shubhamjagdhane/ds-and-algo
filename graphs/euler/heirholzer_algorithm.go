package eulerpath

import "slices"

// reference: https://leetcode.com/problems/reconstruct-itinerary/description/

func EularPathByHierholzerAlgorithm(pairs [][]string) []string {
	// 1. Build a graph from the given pairs
	// 2. Find out the startNode, for some problems startNode is already provided
	// 3. Traverse the graph by visiting all child nodes first then visit the parent

	graph := map[string][]string{}
	indegree := map[string]int{}
	outdegree := map[string]int{}

	// building the graph by using the given points
	for _, pair := range pairs {
		from, to := pair[0], pair[1]
		graph[from] = append(graph[from], to) // assuming the pairs is unidirectional

		// outdegree
		outdegree[from]++
		if _, ok := outdegree[to]; !ok {
			outdegree[to] = 0
		}

		// indegree
		indegree[to]++
		if _, ok := indegree[from]; !ok {
			indegree[from] = 0
		}
	}

	// find out the start and end nodes
	startNode := pairs[0][0]
	// endNode := pairs[0][0]
	for node := range indegree { // either we can traverse indegree or outdegree
		if outdegree[node]-indegree[node] == 1 {
			startNode = node
		}
		// if indegree[node]-outdegree[node] == 1 {
		// 	endNode = node
		// }
	}

	startNode = "JKF" // start node is already provided in the given problem

	// traverse the graph with the start node
	var traverseByDFS func(node string)
	result := []string{}
	traverseByDFS = func(node string) {
		totalChildren := len(graph[node])
		for totalChildren > 0 {
			// note here we are taking child from right of the slice
			// to reduce the time complexity of deletion of node from slice
			// if in some questions ask to consider nodes in lexographical order
			// then we need to sort children in decreasing order
			child := graph[node][totalChildren-1]
			graph[node] = graph[node][:totalChildren-1]
			traverseByDFS(child)
		}
		result = append(result, node)
	}

	traverseByDFS(startNode)

	// as we are traversing children node first and then
	// going back to parent node which means parent node
	// will be on the rightmost of the slice
	// if question ask about to return output from the source node
	// then we need to reverse the result
	slices.Reverse(result)

	return result
}
