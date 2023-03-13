package trim_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
		Problem: https://leetcode.com/problems/trim-a-binary-search-tree/
	  	with each node

		if node.Val == nil => return nil
		because binary tree trait is : left child < val < right child, so:
			if node.Val > high => return left (call again this func) trimBST(node.Left)
			if node.Val < low => return right (call again this func) trimBST(node.Right)

		with left => check left => left = trimBST(node.Left)
		with right => check right => right = trimBST(node.Right)
		return current node

		with each note, is gone to 1 time -> time complexity : O(n)
		we don't use any extra variable -> space complexity : O(1)
*/

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
