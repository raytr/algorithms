package evaluate_division

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	testcases := []struct {
		description string
		equations   [][]string
		values      []float64
		queries     [][]string
		want        []float64
	}{
		{"equations = [[\"a\",\"b\"},{\"b\",\"c\"]], values = [2.0,3.0], queries = [[\"a\",\"c\"},{\"b\",\"a\"},{\"a\",\"e\"},{\"a\",\"a\"},{\"x\",\"x\"]]",
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2, 3},
			[][]string{{"x", "x"}},
			[]float64{-1.00000},
		},
		//{"equations = [[\"a\",\"b\"},{\"b\",\"c\"]], values = [2.0,3.0], queries = [[\"a\",\"c\"},{\"b\",\"a\"},{\"a\",\"e\"},{\"a\",\"a\"},{\"x\",\"x\"]]",
		//	[][]string{{"a", "b"}, {"b", "c"}},
		//	[]float64{2, 3},
		//	[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
		//	[]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		//},

		{"[[\"x1\",\"x2\"},{\"x2\",\"x3\"},{\"x3\",\"x4\"},{\"x4\",\"x5\"]]",
			[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}},
			[]float64{3.0, 4.0, 5.0, 6.0},
			[][]string{{"x2", "x4"}},
			[]float64{20},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			require.ElementsMatch(t, tc.want, calcEquation(tc.equations, tc.values, tc.queries))
		})
	}
}
