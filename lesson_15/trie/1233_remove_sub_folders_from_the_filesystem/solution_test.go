package remove_sub_folders_from_the_filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []string{"/a", "/c/d", "/c/f"}, removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
	assert.Equal(t, []string{"/a"}, removeSubfolders([]string{"/a", "/a/b/c", "/a/b/d"}))
}
