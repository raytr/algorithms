package add_binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		name        string
		a           string
		b           string
		expectation string
	}{
		{"Input: a = \"11\", b = \"1\"", "11", "1", "100"},
		{"a = \"1010\", b = \"1011\"", "1010", "1011", "10101"},
		{"long number",
			"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			"110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, addBinary(tt.a, tt.b))
		})
	}
}
