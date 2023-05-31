package RPNcalculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidExpressions(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		wantFloat  float64
		wantErr    error
	}{
		{"-63 92 + -31 + -60 -", "-63 92 + -31 + -60 -", 58.0, nil},
		{"-31 -92 - 5 * -11 +", "-31 -92 - 5 * -11 +", 294.0, nil},
		{"68 -26 * ewewnot + -83 ^", "68 -26 * not + -83 ^", 0, errors.New("wrong format")},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.expression)
			if err != nil {
				require.Equal(t, tt.wantErr, err)
			} else {
				require.Equal(t, tt.wantFloat, got)
			}
		})
	}
}
