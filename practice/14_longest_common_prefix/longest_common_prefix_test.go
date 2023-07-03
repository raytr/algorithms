package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name        string
		strs        []string
		expectation string
	}{
		{"strs = [\"flower\",\"flow\",\"flight\"]", []string{"flower", "flow", "flight"}, "fl"},
		{"strs = [\"dog\",\"racecar\",\"car\"]", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, longestCommonPrefixBruteForce(tt.strs))
		})
	}
}
