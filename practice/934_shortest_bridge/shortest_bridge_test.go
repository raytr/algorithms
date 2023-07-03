package shortest_bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
}
