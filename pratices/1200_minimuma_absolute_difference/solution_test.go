package minimuma_absolute_difference

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolutionBruteForce(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, minimumAbsDifferenceBruteForce([]int{4, 2, 1, 3}))
}
