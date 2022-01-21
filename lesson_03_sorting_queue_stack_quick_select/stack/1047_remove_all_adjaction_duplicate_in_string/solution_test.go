package remove_all_adjaction_duplicate_in_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "azxzy", removeDuplicates("azxxxzy"))
}
