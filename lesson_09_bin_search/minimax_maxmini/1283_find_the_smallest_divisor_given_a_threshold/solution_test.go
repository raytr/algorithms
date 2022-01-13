package find_the_smallest_divisor_given_a_threshold

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 1, smallestDivisor([]int{21212, 10101, 12121}, 1000000))
}
