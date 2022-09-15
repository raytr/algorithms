package k_th_largest_element_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
