package find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//
//func TestSolution(t *testing.T) {
//	//assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
//	//assert.Equal(t, []int{1, 1}, searchRange([]int{1, 4}, 4))
//	//assert.Equal(t, []int{0, 0}, searchRange([]int{1, 2, 3}, 1))
//	assert.Equal(t, []int{1, 2}, searchRange([]int{1, 2, 2}, 2), []int{1, 2})
//}

func TestSolution(t *testing.T) {
	testCases := []struct {
		name           string
		expression     []int
		expectedResult []int
	}{
		{"happy case", searchRange([]int{1, 2, 2}, 2), []int{1, 2}}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectedResult, tt.expression, tt.name)
		})
	}
}
