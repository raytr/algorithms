package pascals_triangle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{"numRows = 5", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{"numRows = 1", 1, [][]int{{1}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.ElementsMatch(t, tt.want, generate(tt.numRows))
		})
	}
}
