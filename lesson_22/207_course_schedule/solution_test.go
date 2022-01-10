package course_schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, true, canFinish(2, [][]int{{1, 0}}))
	//assert.Equal(t, false, canFinish(2, [][]int{{1, 0}, {0, 1}}))
	assert.Equal(t, true, canFinish(8, [][]int{{1, 0}, {2, 6}, {1, 7}, {6, 4}, {7, 0}, {0, 5}}))
}
