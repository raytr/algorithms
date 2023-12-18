package unique_number_of_occurrences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want bool
	}{
		{"arr = [1,2,2,1,1,3]", []int{1, 2, 2, 1, 1, 3}, true},
		{"arr = [1,2]", []int{1, 2}, false},
		{"arr = [-3,0,1,-3,1,1,1,-3,10,0]", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, uniqueOccurrences(tt.arr))
		})
	}
}
