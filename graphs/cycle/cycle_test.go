package cycle

import "testing"

func TestCycleDFSInUndirectedGraph(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		output     bool
	}{
		{
			name:       "connected undirected graph has cycle",
			totalNodes: 6,
			graph:      [][]int{{1, 2}, {0, 3}, {0, 4}, {1, 5}, {2, 5}, {3, 4}},
			output:     true,
		},
		{
			name:       "connected undirected graph has no cycle",
			totalNodes: 3,
			graph:      [][]int{{1, 2}, {0}, {0}},
			output:     false,
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := isCycleDFSInUndirectedGraph(tt.totalNodes, tt.graph)
			if result != tt.output {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}

func TestCycleBFSInUndirectedGraph(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		output     bool
	}{
		{
			name:       "connected undirected graph has cycle",
			totalNodes: 6,
			graph:      [][]int{{1, 2}, {0, 3}, {0, 4}, {1, 5}, {2, 5}, {3, 4}},
			output:     true,
		},
		{
			name:       "connected undirected graph has no cycle",
			totalNodes: 3,
			graph:      [][]int{{1, 2}, {0}, {0}},
			output:     false,
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := isCycleBFSInUndirectedGraph(tt.totalNodes, tt.graph)
			if result != tt.output {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}

func TestCycleDFSInDirectedGraph(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		output     bool
	}{
		{
			name:       "connected directed graph has cycle",
			totalNodes: 4,
			graph:      [][]int{{1}, {2}, {3}, {0}},
			output:     true,
		},
		{
			name:       "connected directed graph has no cycle",
			totalNodes: 4,
			graph:      [][]int{{1}, {2, 3}, {3}, {}},
			output:     false,
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := isCycleDFSInDirectedGraph(tt.totalNodes, tt.graph)
			if result != tt.output {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}

func TestCycleBFSInDirectedGraph(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		output     bool
	}{
		{
			name:       "connected directed graph has cycle",
			totalNodes: 4,
			graph:      [][]int{{1}, {2}, {3}, {0}},
			output:     true,
		},
		{
			name:       "connected directed graph has no cycle",
			totalNodes: 4,
			graph:      [][]int{{1}, {2, 3}, {3}, {}},
			output:     false,
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := isCycleBFSInDirectedGraph(tt.totalNodes, tt.graph)
			if result != tt.output {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}
