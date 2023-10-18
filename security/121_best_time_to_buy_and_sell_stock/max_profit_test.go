package best_time_to_buy_and_sell_stock

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name        string
		prices      []int
		expectation int
	}{
		{"prices = [7,1,5,3,6,4]", []int{7, 1, 5, 3, 6, 4}, 5},
		{"prices = [7,6,4,3,1]", []int{7, 6, 4, 3, 1}, 0},
		{"prices = [1]", []int{1}, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, maxProfit(tt.prices))
		})
	}
}
