package excel_sheet_column_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTitleToNumber(t *testing.T) {
	testCases := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{"A", "A", 1},
		{"AB", "AB", 28},
		{"ZY", "ZY", 701},
		{"W", "W", 23},
		{"XW", "XW", 647},
		{"FXSHRXW", "FXSHRXW", 2147483647},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, titleToNumber(tt.columnTitle))
		})
	}
}
