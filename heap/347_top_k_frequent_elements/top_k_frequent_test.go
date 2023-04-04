package top_k_frequent_elements

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"nums = [1,1,1,2,2,3], k = 2", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"nums = [1], k = 1", []int{1}, 1, []int{1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotByHeap := topKFrequentByHeap(tt.nums, tt.k)
			gotBySorting := topKFrequentBySorting(tt.nums, tt.k)

			sort.Ints(gotByHeap)
			sort.Ints(gotBySorting)

			resultByHeap := reflect.DeepEqual(tt.want, gotByHeap)
			resultBySorting := reflect.DeepEqual(tt.want, gotBySorting)

			assert.Equal(t, true, resultByHeap)
			assert.Equal(t, true, resultBySorting)
		})
	}
}
