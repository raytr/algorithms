package same_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	p1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	q1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	p2 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	q2 = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	p3 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	q3 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{"p = [1,2,3], q = [1,2,3]", p1, q1, true},
		{"p = [1,2], q = [1,null,2]", p2, q2, false},
		{"p = [1,2,1], q = [1,1,2]", p3, q3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSameTreeRecursion(tt.p, tt.q))
			assert.Equal(t, tt.want, isSameTreeIteration(tt.p, tt.q))
		})
	}
}
