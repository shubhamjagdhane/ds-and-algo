package traversal

import (
	"slices"
	"testing"
)

func TestDFS(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		output     []int
	}{
		{
			name:       "connected undirected graph",
			totalNodes: 6,
			graph:      [][]int{{1, 2}, {0, 3}, {0, 4}, {1, 5}, {2, 5}, {3, 4}},
			output:     []int{0, 1, 3, 5, 4, 2},
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := dfs(tt.totalNodes, tt.graph)
			if !slices.Equal(tt.output, result) {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}

func TestBFS(t *testing.T) {
	testCase := []struct {
		name       string
		graph      [][]int
		totalNodes int
		output     []int
	}{
		{
			name:       "connected undirected graph",
			totalNodes: 6,
			graph:      [][]int{{1, 2}, {0, 3}, {0, 4}, {1, 5}, {2, 5}, {3, 4}},
			output:     []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := bfs(tt.totalNodes, tt.graph)
			if !slices.Equal(tt.output, result) {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}
