package subarray_sums_divisible_by_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysDivByKBruteForce(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"nums = [1,2,3], k = 2", []int{1, 2, 3}, 2, 2},
		{"nums = [4,5,0,-2,-3,1], k = 5", []int{4, 5, 0, -2, -3, 1}, 5, 7},
		{"nums = [5], k = 9", []int{5}, 9, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, subarraysDivByKBruteForce(tt.nums, tt.k))
		})
	}
}

func TestSubarraysDivByK(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"nums = [1,2,3], k = 2", []int{1, 2, 3}, 2, 2},
		{"nums = [4,5,0,-2,-3,1], k = 5", []int{4, 5, 0, -2, -3, 1}, 5, 7},
		{"nums = [5], k = 9", []int{5}, 9, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, subarraysDivByK(tt.nums, tt.k))
		})
	}
}
