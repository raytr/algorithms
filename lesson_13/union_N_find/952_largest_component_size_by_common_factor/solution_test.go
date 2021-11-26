package largest_component_size_by_common_factor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 4, largestComponentSize([]int{4,6,15,35}))
	assert.Equal(t, 2, largestComponentSize([]int{20, 50, 9, 63}))
}
