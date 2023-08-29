package minimum_rounds_to_complete_all_tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumRounds(t *testing.T) {
	testCases := []struct {
		name        string
		tasks       []int
		expectation int
	}{
		{"tasks = [2,2,3,3,2,4,4,4,4,4]", []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4},
		{"tasks = [2,3,3]", []int{2, 3, 3}, -1},
		{"tasks = [5,5,5,5]", []int{5, 5, 5, 5}, 2},
		{"tasks = [5,5,5,5,5]", []int{5, 5, 5, 5, 5}, 2},
		{"tasks = [69,65,62,64,70,68,69,67,60,65,69,62,65,65,61,66,68,61,65,63,60,66,68,66,67,65,63,65,70,69,70,62,68,70,60,68,65,61,64,65,63,62,62,62,67,62,62,61,66,69]", []int{69, 65, 62, 64, 70, 68, 69, 67, 60, 65, 69, 62, 65, 65, 61, 66, 68, 61, 65, 63, 60, 66, 68, 66, 67, 65, 63, 65, 70, 69, 70, 62, 68, 70, 60, 68, 65, 61, 64, 65, 63, 62, 62, 62, 67, 62, 62, 61, 66, 69}, 20},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, minimumRounds(tt.tasks))
		})
	}
}
