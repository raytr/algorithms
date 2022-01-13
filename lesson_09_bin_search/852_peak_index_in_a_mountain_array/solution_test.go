package _52_peak_index_in_a_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 2,peakIndexInMountainArray([]int{3,4,5,1}))
	assert.Equal(t, 1, peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
}
