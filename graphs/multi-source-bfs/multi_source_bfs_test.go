package multisourcebfs

import "testing"

func TestMultiSourceBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "3X3 grid can be rotten",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			expected: 4,
		},
		{
			name: "3X3 grid cannot be rotten",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			expected: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MultiSourceBFS(tc.args.grid)
			if got != tc.expected {
				t.Errorf("expected: %d, got: %d\n", tc.expected, got)
			}
		})
	}
}
