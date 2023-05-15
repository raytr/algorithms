package network_delay_time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name        string
		times       [][]int
		n           int
		k           int
		expectation int
	}{
		{"times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2", [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{"times = [[1,2,1]], n = 2, k = 1", [][]int{{1, 2, 1}}, 2, 1, 1},
		{"times = [[1,2,1]], n = 2, k = 2", [][]int{{1, 2, 1}}, 2, 2, -1},
		{"times = [[1,2,1], [2,1,3]], n = 2, k = 2", [][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2, 3},
		{"times = [[1,2,1], [2,3,2],[1,3,4]], n = 3, k = 1", [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1, 3},
		{"times = [[1,2,1], [2,3,7],[1,3,4],[2,1,2]], n = 3, k = 1", [][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, 3, 1, 4},
		{"times = [[1,2,1], [2,3,7],[1,3,4],[2,1,2]], n = 3, k = 1", [][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, 4, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, networkDelayTime(tt.times, tt.n, tt.k))
		})
	}
}
