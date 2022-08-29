package fruit_into_baskets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, totalFruit([]int{1, 2, 1}))
	assert.Equal(t, 5, totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
	assert.Equal(t, 1, totalFruit([]int{1}))
}
