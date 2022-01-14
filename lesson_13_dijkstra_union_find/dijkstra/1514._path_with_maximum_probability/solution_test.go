package path_with_maximum_probability

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 0.25, maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	assert.Equal(t, 0.3, maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
	assert.Equal(t, 0.21390, maxProbability(5, [][]int{{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3}}, []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04}, 3, 4))
}
