package group_anagrams

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"strs = [eat,tea,tan,ate,nat,bat]", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"strs = []", []string{""}, [][]string{{""}}},
		{"strs = [a]", []string{"a"}, [][]string{{"a"}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := reflect.DeepEqual(groupAnagrams(tt.strs), tt.want)
			assert.Equal(t, true, got)
		})
	}
}
