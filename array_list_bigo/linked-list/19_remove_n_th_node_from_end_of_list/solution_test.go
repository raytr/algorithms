package remove_n_th_node_from_end_of_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input1 = &listNode{
		val: 1,
		next: &listNode{
			val: 2,
			next: &listNode{
				val: 3,
				next: &listNode{
					val: 4,
					next: &listNode{
						val: 5,
					},
				},
			},
		},
	}

	output1 = &listNode{
		val: 1,
		next: &listNode{
			val: 2,
			next: &listNode{
				val: 3,
				next: &listNode{
					val: 5,
				},
			},
		},
	}

	input2 = &listNode{
		val: 2,
	}

	input3 = &listNode{
		val: 1,
		next: &listNode{
			val: 2,
		},
	}
	output3 = &listNode{
		val: 1,
	}

	input4 = &listNode{
		val: 1,
		next: &listNode{
			val: 2,
		},
	}
	output4 = &listNode{
		val: 2,
	}
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		name  string
		input *listNode
		n     int
		want  *listNode
	}{
		{"head = [1,2,3,4,5], n = 2", input1, 2, output1},
		{"head = [1], n = 1", input2, 1, nil},
		{"head = [1,2], n = 1", input3, 1, output3},
		{"head = [1,2], n = 5", input4, 5, output4},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeNthFromEnd(tt.input, tt.n))
		})
	}
}
