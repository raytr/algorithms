package summary_ranges

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSummaryRanges(t *testing.T) {
	testCases := []struct {
		description string
		nums        []int
		want        []string
	}{
		{"nums = [0,1,2,4,5,7]", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"nums = [0,2,3,4,6,8,9]", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{"nums = []", []int{}, []string{}},
		{"nums = [0]", []int{0}, []string{"0"}},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			require.Equal(t, tt.want, summaryRanges(tt.nums))
		})
	}
}
