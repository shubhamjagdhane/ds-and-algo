package dijkstra

import (
	"slices"
	"testing"
)

// fmt.Println(shortestPath(3, [][]int{{0, 1, 1}, {1, 2, 3}, {0, 2, 6}}, 2)) // [4, 3, 0]
// fmt.Println(shortestPath(5, [][]int{{0, 1, 4}, {0, 2, 8}, {1, 4, 6}, {2, 3, 2}, {3, 4, 10}}, 0)) // [0, 4, 8, 10, 10]
//

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
			name: "undirected connected graph",
			args: args{
				n:     3,
				edges: [][]int{{0, 1, 3}, {1, 2, 2}, {0, 2, 6}},
				src:   2,
			},
			expected: []int{5, 2, 0},
		},

		{
			name: "undirected connected graph",
			args: args{
				n:     6,
				edges: [][]int{{0, 1, 4}, {0, 2, 8}, {1, 2, 3}, {1, 4, 6}, {2, 3, 1}, {3, 5, 6}, {4, 5, 7}},
				src:   0,
			},
			expected: []int{0, 4, 7, 8, 10, 14},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := shortestPath(tc.args.n, tc.args.edges, tc.args.src)
			if !slices.Equal(tc.expected, got) {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
