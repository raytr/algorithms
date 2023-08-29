package pairs_of_songs_with_total_durations_divisible_by_60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	testCases := []struct {
		name        string
		time        []int
		expectation int
	}{
		{"time = [30,20,150,100,40]", []int{30, 20, 150, 100, 40}, 3},
		{"time = [60,60,60]", []int{60, 60, 60}, 3},
		{"time = [60]", []int{60}, 1},
		{"time = [50]", []int{50}, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, numPairsDivisibleBy60(tt.time))
		})
	}
}
