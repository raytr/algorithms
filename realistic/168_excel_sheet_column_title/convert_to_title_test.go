package excel_sheet_column_title

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name         string
		columnNumber int
		want         string
	}{
		{"1", 1, "A"},
		{"28", 28, "AB"},
		{"52", 52, "AZ"},
		{"701", 701, "ZY"},
		{"730", 730, "ABB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, convertToTitle(tt.columnNumber))
		})
	}
}
