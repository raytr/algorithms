package optimize_water_distribution_in_a_village

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 131704, minCostToSupplyWater(5, []int{46012, 72474, 64965, 751, 33304}, [][]int{{2, 1, 6719}, {3, 2, 75312}, {5, 3, 44918}}))
}
