package two_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		description string
		nums        []int
		target      int
		expectation []int
	}{
		{"nums = [2,7,11,15], target = 9", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"nums = [3,2,4], target = 6", []int{3, 2, 4}, 6, []int{1, 2}},
		{"nums = [3,3], target = 6", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			require.ElementsMatch(t, testCase.expectation, twoSum(testCase.nums, testCase.target))
		})
	}
}
