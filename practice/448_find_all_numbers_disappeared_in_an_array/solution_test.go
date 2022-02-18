package find_all_numbers_disappeared_in_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int{1, 4}, findDisappearedNumbers([]int{2, 3, 2, 3}))
}
