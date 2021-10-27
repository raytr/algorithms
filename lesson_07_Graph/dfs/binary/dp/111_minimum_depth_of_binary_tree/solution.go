package minimum_depth_of_binary_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	problem: https://leetcode.com/problems/minimum-depth-of-binary-tree/

	we have minDepth(root *TreeNode) int
	parent's height = min(left child's height, min child's height) + 1
	=> define dfs:
		- node.Val == nil return +oo
		- node.Left == nil & node.Right == nil => return 1
		- min(dfs(node.left), dfs(node.right)) + 1

	complexity: O(2n)
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return int(math.Pow(10, 5))
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return min(dfs(node.Left), dfs(node.Right)) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
