package connecting_cities_with_minimum_cost

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 6, minimumCost(3, [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}))
}
