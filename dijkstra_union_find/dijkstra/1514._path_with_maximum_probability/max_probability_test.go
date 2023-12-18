package path_with_maximum_probability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProbability(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
		want     float64
	}{
		{"Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2, 0.25},
		{"Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2, 0.3},
		{"Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2", 5, [][]int{{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3}}, []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04}, 3, 4, 0.21390},
		{"Input: n = 500, edges = [[193,229],[133,212],[224,465]], succProb = [0.91,0.78,0.64], start = 4, end = 364", 500, [][]int{{193, 229}, {133, 212}, {224, 465}}, []float64{0.91, 0.78, 0.64}, 3, 364, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProbability(tt.n, tt.edges, tt.succProb, tt.start, tt.end))
		})
	}
}
