package binary_tree_inorder_traversal

/*
https://leetcode.com/problems/binary-tree-inorder-traversal/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	if root == nil {
		return result
	}

	dfs(root)

	return result
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// first recursive on left subtree
	dfs(root.Left)

	// appent to result
	result = append(result, root.Val)

	//then recursive on right subtree
	dfs(root.Right)
}
