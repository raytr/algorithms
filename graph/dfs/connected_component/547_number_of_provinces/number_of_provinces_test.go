package number_of_provinces

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberOfProvinces(t *testing.T) {
	testCases := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{"isConnected = [[1,1,0],[1,1,0],[0,0,1]]", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{"isConnected = [[1,0,0],[0,1,0],[0,0,1]]", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, findCircleNum(tt.isConnected))
		})
	}
}
