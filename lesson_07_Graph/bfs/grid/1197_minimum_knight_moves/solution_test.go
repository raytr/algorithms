package minimum_knight_moves

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 1, minKnightMoves(2,1))
	//assert.Equal(t, 4, minKnightMoves(5,5))
	assert.Equal(t, 58, minKnightMoves(-87, 83))
}
