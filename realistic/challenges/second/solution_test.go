package second

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTheLargestNumber(t *testing.T) {
	testCases := []struct {
		name string
		N    int
		want int
	}{
		{"4268", 4268, 54268},
		{"6750", 6750, 67550},
		{"-999", -999, -5999},
		{"-394", -394, -3594},
		{"5268", 5268, 55268},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, TheLargestNumber(tt.N))
		})
	}
}
