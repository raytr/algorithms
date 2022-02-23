package word_ladder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	a := ladderLength("hit", "hot", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	b := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(a, b)

	assert.Equal(t, 2, ladderLength("hit", "hot", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	assert.Equal(t, 5, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
