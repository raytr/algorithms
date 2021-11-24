package satisfiability_of_equality_equations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, false, equationsPossible([]string{"a==b", "b!=a"}))
	assert.Equal(t, true, equationsPossible([]string{"b==a", "a==b"}))
	assert.Equal(t, true, equationsPossible([]string{"a==b", "b==c", "a==c"}))
	assert.Equal(t, false, equationsPossible([]string{"a==b", "b!=c", "c==a"}))
	assert.Equal(t, true, equationsPossible([]string{"c==c", "b==d", "x!=z"}))
}
