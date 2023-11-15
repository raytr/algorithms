package invert_binary_tree

/*
	problem: https://leetcode.com/problems/invert-binary-tree
	using recursive, newNode.Left = invertTree(Node.Right), newNode.Right = invertTree(Node.Left)

	Since each node in the tree is visited only once, therefor, time complexity: O(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := &TreeNode{
		Val: root.Val,
	}

	res.Left = invertTree(root.Right)
	res.Right = invertTree(root.Left)

	return res
}
