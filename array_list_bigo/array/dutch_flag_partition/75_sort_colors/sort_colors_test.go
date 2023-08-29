package sort_colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name        string
		nums        []int
		expectation []int
	}{
		{"nums = [2,0,2,1,1,0]", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"nums = [2,0,1]", []int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.expectation, sortColors(tt.nums))
		})
	}
}
