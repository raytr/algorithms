package max_sum_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSumArray(t *testing.T) {
	//assert.Equal(t, 6, maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}), "wrong at first")
	//assert.Equal(t, 1, maxSubArray([]int{1}), "wrong at second")
	//assert.Equal(t, 23, maxSubArray([]int{5,4,-1,7,8}), "wrong at third")
	//assert.Equal(t, 1, maxSubArray([]int{-2, 1}), "wrong at forth")
	assert.Equal(t, -1, maxSubArray([]int{-2, -1}), "wrong at fifth")
}
