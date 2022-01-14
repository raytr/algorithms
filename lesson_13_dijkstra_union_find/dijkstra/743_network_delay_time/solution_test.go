package network_delay_time

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	assert.Equal(t, -1, networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
	assert.Equal(t, 3, networkDelayTime([][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2))
	assert.Equal(t, 3, networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1))
	assert.Equal(t, 4, networkDelayTime([][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, 3, 1))
	assert.Equal(t, -1, networkDelayTime([][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, 4, 1))
}
