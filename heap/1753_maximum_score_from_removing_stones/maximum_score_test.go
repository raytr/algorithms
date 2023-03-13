package maximum_score_from_removing_stones

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{"(2,4,6)", 2, 4, 6, 6},
		{"(4,4,6)", 4, 4, 6, 7},
		{"(1,8,8)", 1, 8, 8, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maximumScore(tt.a, tt.b, tt.c))
		})
	}
}
