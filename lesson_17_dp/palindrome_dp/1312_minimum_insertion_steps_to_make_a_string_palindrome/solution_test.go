package minimum_insertion_steps_to_make_a_string_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 0, countSubstrings("zzazz"))
	assert.Equal(t, 2, minInsertions("mbadm"))
	//assert.Equal(t, 5, countSubstrings("leetcode"))

}
