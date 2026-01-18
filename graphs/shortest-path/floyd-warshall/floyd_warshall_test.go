package flyodwarshall

import (
	"math"
	"slices"
	"testing"
)

func TestShortestPath(t *testing.T) {

	/*
		edges: [][]int{{0, 1, 3}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {4, 3, 5}, {3, 5, 4}},
	*/
	type args struct {
		n     int
		graph [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected [][]int
	}{
		{
			name: "directed graph",
			args: args{
				n: 5,
				graph: [][]int{
					{0, 4, math.MaxInt, 5, math.MaxInt},
					{math.MaxInt, 0, 1, math.MaxInt, 6},
					{2, math.MaxInt, 0, 3, math.MaxInt},
					{math.MaxInt, math.MaxInt, 1, 0, 2},
					{1, math.MaxInt, math.MaxInt, 4, 0},
				},
			},
			expected: [][]int{
				{0, 4, 5, 5, 7},
				{3, 0, 1, 4, 6},
				{2, 6, 0, 3, 5},
				{3, 7, 1, 0, 2},
				{1, 5, 5, 4, 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ShortestPath(tc.args.n, tc.args.graph)
			if !verify(tc.expected, output) {
				t.Errorf("Expected %v, got %v", tc.expected, output)
			}
		})
	}
}

func verify(slice1, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for _, s1 := range slice1 {
		isPresent := false
		for _, s2 := range slice2 {
			if slices.Equal(s1, s2) {
				isPresent = true
				break
			}
		}
		if !isPresent {
			return false
		}
	}
	return true
}
