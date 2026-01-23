package disjointsetunion

import "testing"

func TestFind(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		parent   []int
		expected int
	}{
		{
			name:     "parent of the node is itself",
			parent:   []int{0, 1, 2, 3},
			x:        3,
			expected: 3,
		},
		{
			name:     "node has a differnt parent than self",
			parent:   []int{0, 1, 0, 3},
			x:        2,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Find(tc.x, tc.parent)
			if got != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, got)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		x      int
		y      int
		parent []int
		rank   []int
	}
	type expected struct {
		xParent int
		rankOfY int
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "union 1 with 2, 2 should be new parent of 1",
			args: args{
				x:      1,
				y:      2,
				parent: []int{0, 1, 2, 3},
				rank:   []int{0, 0, 0, 0},
			},
			expected: expected{
				xParent: 2,
				rankOfY: 1,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			Union(tc.args.x, tc.args.y, tc.args.parent, tc.args.rank)

			xParent := Find(tc.args.x, tc.args.parent)

			if xParent != tc.expected.xParent || tc.args.rank[tc.args.y] != tc.expected.rankOfY {
				t.Errorf("expected x & y parents should be same, received xParent:%d and yParent: %d as well as rank of y should be %d, received: %d", xParent, tc.expected.xParent, tc.expected.rankOfY, tc.args.rank[tc.args.y])
			}
		})
	}
}
