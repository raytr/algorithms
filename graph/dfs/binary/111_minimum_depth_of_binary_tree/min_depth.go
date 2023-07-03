package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	problem: https://leetcode.com/problems/minimum-depth-of-binary-tree/

	we have minDepth(root *TreeNode) int
	parent's height = getMin(left child's height, getMin child's height) + 1
	=> define dfs:
		- node.val == nil return +oo
		- node.Left == nil & node.Right == nil => return 1
		- getMin(dfs(node.left), dfs(node.right)) + 1

	complexity: O(2n)
*/

var maxInt = 100000

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return maxInt
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return getMin(dfs(node.Left), dfs(node.Right)) + 1
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
