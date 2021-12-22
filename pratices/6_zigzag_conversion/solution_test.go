package zigzag_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "AB", convert("AB", 1))
}
