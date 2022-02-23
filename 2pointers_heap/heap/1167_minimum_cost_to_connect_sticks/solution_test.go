package minimum_cost_to_connect_sticks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 14, connectSticks([]int{2, 4, 3}))
	assert.Equal(t, 30, connectSticks([]int{1, 8, 3, 5}))
}
