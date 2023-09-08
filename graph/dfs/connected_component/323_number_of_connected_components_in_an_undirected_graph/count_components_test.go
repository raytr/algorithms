package number_of_connected_components_in_an_undirected_graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountComponents(t *testing.T) {
	testCases := []struct {
		name        string
		n           int
		edges       [][]int
		expectation int
	}{
		{"n = 5, edges = [[0,1},{1,2},{3,4]]", 5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{"n = 5, edges = [[0,1},{1,2},{2,3},{3,4]]", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, countComponents(tt.n, tt.edges))
		})
	}
}
