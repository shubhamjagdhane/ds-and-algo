package stronglyconnectedcomponants

import "testing"

func TestScc(t *testing.T) {
	type args struct {
		graph [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "connected directed graph with 5 vertices",
			args: args{
				graph: [][]int{{2, 3}, {0}, {1}, {4}, {}},
			},
			expected: 3,
		},
		{
			name: "connected directed graph with 11 vertices",
			args: args{
				graph: [][]int{{1}, {2, 3}, {0}, {4}, {5}, {3}, {5, 7}, {8}, {9}, {6, 10}, {}},
			},
			expected: 4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := KosarajusAlgorithm(tc.args.graph)
			if tc.expected != got {
				t.Errorf("expected: %d, got: %d", tc.expected, got)
			}
		})
	}
}
