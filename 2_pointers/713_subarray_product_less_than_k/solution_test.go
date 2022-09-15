package subarray_product_less_than_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 8, numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	assert.Equal(t, 0, numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
	assert.Equal(t, 1, numSubarrayProductLessThanK([]int{1}, 2))
	assert.Equal(t, 0, numSubarrayProductLessThanK([]int{1}, 0))
	assert.Equal(t, 3, numSubarrayProductLessThanK([]int{0, 1}, 4))
	assert.Equal(t, 18, numSubarrayProductLessThanK([]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19))
	assert.Equal(t, 1, numSubarrayProductLessThanK([]int{76, 8, 75, 22}, 18))
}
