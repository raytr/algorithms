package minimum_limit_of_balls_in_a_bag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 7, minimumSize([]int{7, 17}, 2))
	assert.Equal(t, 2, minimumSize([]int{2, 4, 8, 2}, 4))
	assert.Equal(t, 9, minimumSize([]int{9, 18}, 1))
}
