package binary_tree_right_side_view

//problem : https://leetcode.com/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		canSee := q[0].Val
		for _, n := range q {
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			canSee = n.Val
			q = q[1:]
		}
		res = append(res, canSee)
	}
	return res
}
