package min_cost_to_connect_all_points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 20, minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
	assert.Equal(t, 18, minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))
	assert.Equal(t, 4, minCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}))
	assert.Equal(t, 4000000, minCostConnectPoints([][]int{{-1000000, -1000000}, {1000000, 1000000}}))
	assert.Equal(t, 0, minCostConnectPoints([][]int{{0, 0}}))
}
