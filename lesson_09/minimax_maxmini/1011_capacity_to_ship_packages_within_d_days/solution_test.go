package capacity_to_ship_packages_within_d_days

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 15, shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
