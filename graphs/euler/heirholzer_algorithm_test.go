package eulerpath

import "testing"

func TestHeirholzerAlgorithm(t *testing.T) {
	type args struct {
		pairs [][]string
	}
	testCases := []struct {
		name   string
		args   args
		output []string
	}{
		{
			name: "graph with five nodes and four edges",
			args: args{
				pairs: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			},
			output: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name: "graph with three nodes and five edges",
			args: args{
				pairs: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			},
			output: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}
