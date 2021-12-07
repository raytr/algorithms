package k_th_largest_element_in_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
