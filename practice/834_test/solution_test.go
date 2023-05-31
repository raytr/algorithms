package test_834

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfDistancesInTree(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		edges       [][]int
		expectation []int
	}{
		{"n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]", 6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, []int{8, 12, 6, 10, 10, 10}},
		{"n = 1, edges = []", 1, [][]int{}, []int{0}},
		{"n = 2, edges = [[1,0]]", 2, [][]int{{1, 0}}, []int{1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reflect.DeepEqual(sumOfDistancesInTree(tt.n, tt.edges), tt.expectation)
			assert.Equal(t, true, got)
		})
	}
}
