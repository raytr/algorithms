package maximum_value_after_insertion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		param1 string
		param2 int
		want   string
	}{
		{"0", 5, "50"},
		{"7643", 5, "76543"},
		{"999", 5, "9995"},
		{"4321", 5, "54321"},
		{"-999", 5, "-5999"},
		{"-222", 5, "-2225"},
		{"-3467", 5, "-34567"},
	}

	for _, tt := range testCases {
		t.Run(tt.want, func(t *testing.T) {
			require.Equal(t, tt.want, maxValue(tt.param1, tt.param2))
		})
	}
}
