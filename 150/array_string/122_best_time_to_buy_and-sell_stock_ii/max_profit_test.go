package best_time_to_buy_and_sell_stock_ii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		description string
		prices      []int
		expectation int
	}{
		{"prices = [7,1,5,3,6,4]", []int{7, 1, 5, 3, 6, 4}, 7},
		{"prices = [1,2,3,4,5]", []int{1, 2, 3, 4, 5}, 4},
		{"prices = [7,6,4,3,1]", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			require.Equal(t, tt.expectation, maxProfit(tt.prices))
		})
	}
}
