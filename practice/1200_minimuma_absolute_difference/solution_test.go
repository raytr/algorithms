package minimuma_absolute_difference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutionBruteForce(t *testing.T) {
	//assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, minimumAbsDifferenceBruteForce([]int{4, 2, 1, 3}))
	assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, minimumAbsDifferenceSortingFirst([]int{4, 2, 1, 3}))
}
