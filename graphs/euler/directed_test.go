package eulerpath

import "testing"

func TestIsEulerCircuitInDirectedGraph(t *testing.T) {
	type args struct {
		n     int
		graph [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "graph having Euler Circuit",
			args: args{
				n:     4,
				graph: [][]int{{1, 2}, {0, 3}, {1}, {0}},
			},
			expected: 2,
		},
		{
			name: "graph having Euler Path",
			args: args{
				n:     5,
				graph: [][]int{{1}, {2}, {3}, {4}, {2}},
			},
			expected: 1,
		},
		{
			name: "graph neither has Euler Circuit nor Euler Path",
			args: args{
				n:     3,
				graph: [][]int{{}, {0}, {0}},
			},
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsEulerCircuitInDirectedGraph(tc.args.n, tc.args.graph)
			if tc.expected != got {
				t.Errorf("Expected: %d, got %d\n", tc.expected, got)
			}
		})
	}
}
