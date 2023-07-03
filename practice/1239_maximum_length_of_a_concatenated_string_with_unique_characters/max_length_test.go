package maximum_length_of_a_concatenated_string_with_unique_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength(t *testing.T) {
	testCases := []struct {
		name        string
		input       []string
		expectation int
	}{
		{"arr = [\"un\",\"iq\",\"ue\"]", []string{"un", "iq", "ue"}, 4},
		{"arr = [\"uu\",\"ii\",\"oe\"]", []string{"uu", "ii", "oe"}, 2},
		{"arr = [\"cha\",\"r\",\"act\",\"ers\"", []string{"cha", "r", "act", "ers"}, 6},
		{"arr = [\"abcdefghijklmnopqrstuvwxyz\"]", []string{"abcdefghijklmnopqrstuvwxyz"}, 26},
		{"arr = [\"abcdefghijklmnopqrstuvwxyzz\"]", []string{"abcdefghijklmnvwxyzz"}, 0},
		{"arr = [\"jnfbyktlrqumowxd\",\"mvhgcpxnjzrdei\"]", []string{"jnfbyktlrqumowxd", "mvhgcpxnjzrdei"}, 16},
		{"arr = [\"ab\",\"cd\",\"cde\",\"cdef\",\"efg\",\"fgh\",\"abxyz\"], expectation=abxyzcdefcd", []string{"ab", "cd", "cde", "cdef", "efg", "fgh", "abxyz"}, 11},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, maxLength(tt.input))
		})
	}
}
