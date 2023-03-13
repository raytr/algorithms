package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"1,1,2", []int{1, 1, 2}, 2},
		{"0,0,1,1,1,2,2,3,3,4", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
