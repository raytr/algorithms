package run_sum_of_1day_array

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumOf1DayArray(t *testing.T) {

	testCases := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
		{[]int{1}, []int{1}},
	}

	for _, testCase := range testCases {
		result := runningSum(testCase.input)
		require.Equal(t, result, testCase.want, fmt.Sprintf("wrong at %v", testCase.input))
	}

}
