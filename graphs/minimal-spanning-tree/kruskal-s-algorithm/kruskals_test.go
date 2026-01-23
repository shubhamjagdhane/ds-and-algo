package kruskalsalgorithm

import "testing"

func TestMinimalSpanningTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "connected undirected graph with 3 vertices",
			args: args{
				n:     3,
				edges: [][]int{{0, 1, 5}, {1, 2, 3}, {0, 2, 1}},
			},
			expected: 4,
		},
		{
			name: "connected undirected graph with 7 vertices",
			args: args{
				n:     7,
				edges: [][]int{{0, 1, 5}, {0, 3, 20}, {1, 2, 5}, {3, 4, 1}, {2, 3, 5}, {4, 5, 2}, {4, 6, 4}, {5, 6, 2}},
			},
			expected: 20,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minimumSpanningTree(tc.args.n, tc.args.edges)
			if tc.expected != got {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
