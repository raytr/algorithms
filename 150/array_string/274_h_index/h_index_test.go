package h_index

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHIndex(t *testing.T) {
	testcases := []struct {
		description string
		citations   []int
		want        int
	}{
		{"citations = [3,0,6,1,5]", []int{3, 0, 6, 1, 5}, 3},
		{"citations = [1,3,1]", []int{1, 3, 1}, 1},
		{"citations = [1]", []int{1}, 1},
		{"citations = [100]", []int{100}, 1},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, hIndex(tc.citations))
		})
	}
}
