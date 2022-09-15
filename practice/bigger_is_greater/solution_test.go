package bigger_is_greater

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//func TestSolution(t *testing.T) {
//	assert.Equal(t, "hegf", biggerIsGreater("hefg"))
//	assert.Equal(t, "dhkc", biggerIsGreater("dhck"))
//	assert.Equal(t, "hcdk", biggerIsGreater("dkhc"))
//	assert.Equal(t, "abdc", biggerIsGreater("abcd"))
//	assert.Equal(t, "lmon", biggerIsGreater("lmno"))
//	assert.Equal(t, "no answer", biggerIsGreater("bb"))
//	assert.Equal(t, "no answer", biggerIsGreater("dcba"))
//	assert.Equal(t, "no answer", biggerIsGreater("dcbb"))
//	assert.Equal(t, "acbd", biggerIsGreater("abdc"))
//	assert.Equal(t, "abdc", biggerIsGreater("abcd"))
//	assert.Equal(t, "fedcbabdc", biggerIsGreater("fedcbabcd"))
//}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		w    string
		want string
	}{
		{"hefg", "hefg", "hegf"},
		{"dhck", "dhck", "dhkc"},
		{"dkhc", "dkhc", "hcdk"},
		{"abcd", "abcd", "abdc"},
		{"lmno", "lmno", "lmon"},
		{"bb", "bb", "no answer"},
		{"dcba", "dcba", "no answer"},
		{"dcbb", "dcbb", "no answer"},
		{"abdc", "abdc", "acbd"},
		{"abcd", "abcd", "abdc"},
		{"fedcbabcd", "fedcbabcd", "fedcbabdc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, biggerIsGreater(tt.w))
		})
	}
}
