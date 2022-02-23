package maximum_depth_of_binary_tree

// problem: https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(dfs(node.Left), dfs(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
