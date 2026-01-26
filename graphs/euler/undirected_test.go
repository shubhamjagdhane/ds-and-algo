package eulerpath

import "testing"

func TestIsEulerCircuitInUndirectedGraph(t *testing.T) {
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
				n:     3,
				graph: [][]int{{1, 2}, {0, 2}, {0, 1}},
			},
			expected: 2,
		},
		{
			name: "graph having Euler Path",
			args: args{
				n:     3,
				graph: [][]int{{1, 2}, {0}, {0}},
			},
			expected: 1,
		},
		{
			name: "graph neither has Euler Circuit nor Euler Path",
			args: args{
				n:     5,
				graph: [][]int{{1, 2}, {0}, {0}, {4}, {3}},
			},
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsEulerCircuitInUndirectedGraph(tc.args.n, tc.args.graph)
			if tc.expected != got {
				t.Errorf("Expected: %d, got %d\n", tc.expected, got)
			}
		})
	}
}
