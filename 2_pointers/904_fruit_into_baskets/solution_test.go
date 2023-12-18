package fruit_into_baskets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalFruit(t *testing.T) {
	testCases := []struct {
		name   string
		fruits []int
		want   int
	}{
		{"fruits = [1,2,1]", []int{1, 2, 1}, 3},
		{"fruits = [0, 1, 2, 1, 2, 0, 1, 2]", []int{0, 1, 2, 1, 2, 0, 1, 2}, 4},
		{"fruits = [0,1,2,2]", []int{0, 1, 2, 2}, 3},
		{"fruits = [1,1,2,3,2,3,2,3,2,3]", []int{1, 1, 2, 3, 2, 3, 2, 3, 2, 3}, 8},
		{"fruits = [1,2,3,2,2]", []int{1, 2, 3, 2, 2}, 4},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, totalFruit(tt.fruits))
		})
	}
}
