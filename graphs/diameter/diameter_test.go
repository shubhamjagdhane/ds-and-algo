package diameter

import "testing"

func TestDiameterInUnidirectedGraph(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "graph with 3 node and 2 edges",
			args: args{
				pairs: [][]int{{0, 1}, {0, 2}},
			},
			expected: 2,
		},
		{
			name: "graph with 6 node and 5 edges",
			args: args{
				pairs: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 5}},
			},
			expected: 4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := DiameterInUndirectedGraph(tc.args.pairs)
			if tc.expected != got {
				t.Errorf("expected: %d, got: %d\n", tc.expected, got)
			}

		})
	}
}
