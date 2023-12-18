package run_sum_of_1day_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumOf1DayArray(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{"nums = [1,2,3,4]", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"nums = [1, 1, 1, 1, 1]", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"nums = [13, 1, 2, 10, 1]", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
		{"nums = [1]", []int{1}, []int{1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, runningSum(tt.nums))
		})
	}

}
