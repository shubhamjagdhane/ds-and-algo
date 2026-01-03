package topologicalsort

import (
	"slices"
	"testing"
)

func TestTopologicalSortByDFS(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		outputs    [][]int
	}{
		{
			name:       "Directed Acyclic Graph with one vertex having no incoming edge",
			totalNodes: 5,
			graph:      [][]int{{1}, {2}, {3}, {4}, {}},
			outputs:    [][]int{{0, 1, 2, 3, 4}},
		},

		{
			name:       "Directed Acyclic Graph with two vertices having no incoming edge",
			totalNodes: 6,
			graph:      [][]int{{}, {0}, {3}, {}, {1, 2}, {0, 3}},
			outputs:    [][]int{{5, 4, 2, 3, 1, 0}, {4, 5, 1, 2, 0, 3}},
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := TopologiclSortByDFS(tt.totalNodes, tt.graph)
			isValid := false
			for _, output := range tt.outputs {
				if slices.Equal(result, output) {
					isValid = true
				}
			}
			if !isValid {
				t.Errorf("Expected either one of %v, got %v", tt.outputs, result)
			}
		})
	}
}

func TestTopologicalSortByBFS(t *testing.T) {
	testCase := []struct {
		name       string
		totalNodes int
		graph      [][]int
		outputs    [][]int
	}{
		{
			name:       "Directed Acyclic Graph with one vertex having no incoming edge",
			totalNodes: 5,
			graph:      [][]int{{1}, {2}, {3}, {4}, {}},
			outputs:    [][]int{{0, 1, 2, 3, 4}},
		},

		{
			name:       "Directed Acyclic Graph with two vertices having no incoming edge",
			totalNodes: 6,
			graph:      [][]int{{}, {0}, {3}, {}, {1, 2}, {0, 3}},
			outputs:    [][]int{{5, 4, 2, 3, 1, 0}, {4, 5, 1, 2, 0, 3}},
		},
		/*
		* more test cases
		 */
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result := TopologiclSortByBFS(tt.totalNodes, tt.graph)
			isValid := false
			for _, output := range tt.outputs {
				if slices.Equal(result, output) {
					isValid = true
				}
			}
			if !isValid {
				t.Errorf("Expected either one of %v, got %v", tt.outputs, result)
			}
		})
	}
}
