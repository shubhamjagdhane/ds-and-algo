package bellmanford

import (
	"slices"
	"testing"
)

func TestShortestPath(t *testing.T) {
	type args struct {
		n     int     // total number of vertices
		edges [][]int // total number of edges betwen vertices
		src   int     // source node from which to find minimal distance of every node
	}
	testCases := []struct {
		name     string
		args     args
		expected []int
	}{

		{
			name: "undirected connected graph (test case taken from Dijkstra's test)",
			args: args{
				n:     3,
				edges: [][]int{{0, 1, 3}, {1, 0, 3}, {1, 2, 2}, {2, 1, 2}, {0, 2, 6}, {2, 0, 6}},
				src:   2,
			},
			expected: []int{5, 2, 0},
		},

		{
			name: "undirected connected graph (test case taken from Dijkstra's test)",
			args: args{
				n:     6,
				edges: [][]int{{0, 1, 4}, {1, 0, 4}, {0, 2, 8}, {2, 0, 8}, {1, 2, 3}, {2, 1, 3}, {1, 4, 6}, {4, 1, 6}, {2, 3, 1}, {3, 2, 1}, {3, 5, 6}, {5, 3, 6}, {4, 5, 7}, {4, 5, 7}},
				src:   0,
			},
			expected: []int{0, 4, 7, 8, 10, 14},
		},
		{
			name: "directed connected graph (test case taken from Dijkstra's test)",
			args: args{
				n:     6,
				edges: [][]int{{0, 1, 3}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {4, 3, 5}, {3, 5, 4}},
				src:   0,
			},
			expected: []int{0, 3, 6, 7, 5, 11},
		},
		{
			name: "directed connected with -ve edges graph",
			args: args{
				n:     6,
				edges: [][]int{{0, 1, 3}, {1, 2, 3}, {2, 3, 1}, {1, 4, 2}, {4, 3, -6}, {3, 5, -2}},
				src:   0,
			},
			expected: []int{0, 3, 6, -1, 5, -3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := shortestPath(tc.args.n, tc.args.edges, tc.args.src)
			if !slices.Equal(tc.expected, got) {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
