package redundant_connection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int{4, 5}, findRedundantConnection([][]int{{1, 5}, {3, 4}, {3, 5}, {4, 5}, {2, 4}}))
}
