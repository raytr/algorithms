package product_of_array_except_self

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testcases := []struct {
		description string
		nums        []int
		expectation []int
	}{
		{"nums = [1,2,3,4]", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"nums = [-1,1,0,-3,3]", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			require.ElementsMatch(t, testcase.expectation, productExceptSelf(testcase.nums))
		})
	}
}
