package network

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//func TestSolution(t *testing.T) {
//	assert.Equal(t, []int32{1, 3, 2, 3, 0, 0, 0, 0, 0}, distanceCount([]int32{9, 1, 4, 9, 0, 4, 8, 9, 0, 1}))
//	//assert.Equal(t, []int32{1}, distanceCount([]int32{9, 1}))
//}

func TestDistanceCount(t *testing.T) {
	testCases := []struct {
		name string
		T    []int32
		want []int32
	}{
		{"[9, 1, 4, 9, 0, 4, 8, 9, 0, 1]", []int32{9, 1, 4, 9, 0, 4, 8, 9, 0, 1}, []int32{1, 3, 2, 3, 0, 0, 0, 0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, distanceCount(tt.T))
		})
	}
}
