package search_in_a_binary_search_tree

// problem: https://leetcode.com/problems/search-in-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
