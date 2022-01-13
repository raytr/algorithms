package bigger_is_greater

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "hegf", biggerIsGreater("hefg"))
	assert.Equal(t, "dhkc", biggerIsGreater("dhck"))
	assert.Equal(t, "hcdk", biggerIsGreater("dkhc"))
	assert.Equal(t, "abdc", biggerIsGreater("abcd"))
	assert.Equal(t, "lmon", biggerIsGreater("lmno"))
	assert.Equal(t, "no answer", biggerIsGreater("bb"))
	assert.Equal(t, "no answer", biggerIsGreater("dcba"))
	assert.Equal(t, "no answer", biggerIsGreater("dcbb"))
	assert.Equal(t, "acbd", biggerIsGreater("abdc"))
	assert.Equal(t, "abdc", biggerIsGreater("abcd"))
	assert.Equal(t, "fedcbabdc", biggerIsGreater("fedcbabcd"))
}
