package kth_smallest_element_in_a_bst

/*

	problem: https://leetcode.com/problems/kth-smallest-element-in-a-bst/

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var smallestTh, res int

func kthSmallest(root *TreeNode, k int) int {
	smallestTh, res = 0, 0
	dfs(root, k)
	return res
}

func dfs(root *TreeNode, k int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		dfs(root.Left, k)
	}
	smallestTh++
	if smallestTh == k {
		res = root.Val
		return
	}
	if root.Right != nil {
		dfs(root.Right, k)
	}
}
