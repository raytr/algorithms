package subarrays_with_k_different_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 7, subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	assert.Equal(t, 3, subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}
