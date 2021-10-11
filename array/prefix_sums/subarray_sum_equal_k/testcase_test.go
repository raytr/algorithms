package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseTest(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{5, -1, 3, 2, 9, 4}, 5), "wrong at first")
	assert.Equal(t, 3, subarraySum([]int{2, 5, 0, 2, 5}, 5), "wrong at second")
	assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2), "wrong at third")
	assert.Equal(t, 4, subarraySum([]int{1, 2, 1, 2, 1}, 3), "wrong at forth")
	assert.Equal(t, 1, subarraySum([]int{-1, -1, 1}, 0), "wrong at fifth")
	assert.Equal(t, 3, subarraySum([]int{1, -1, 0}, 0), "wrong at sixth")
	assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 3), "wrong at seventh")

}
