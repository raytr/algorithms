package buy_two_chocolates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyChoco(t *testing.T) {
	testCases := []struct {
		name   string
		prices []int
		money  int
		want   int
	}{
		{"prices = [1,2,2], money = 3", []int{1, 2, 2}, 3, 0},
		{"prices = [1,1], money = 1", []int{1, 1}, 1, 1},
		{"prices = [3,2,3], money = 3", []int{3, 2, 3}, 3, 3},
		{"prices = [3,2,3], money = 3", []int{3, 2, 3}, 3, 3},
		{"prices = [7], money = 3", []int{7}, 3, 3},
		{"prices = [3], money = 7", []int{3}, 7, 7},
		{"prices = [41,1,28,2,92,97,1,87], money = 68", []int{41, 1, 28, 2, 92, 97, 1, 87}, 68, 66},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, buyChoco(tt.prices, tt.money))
		})
	}
}
