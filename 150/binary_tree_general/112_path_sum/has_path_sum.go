package path_sum

/*
	problem: https://leetcode.com/problems/path-sum


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return findLeaf(root, 0, targetSum)
}

func findLeaf(root *TreeNode, curSum, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		if curSum+root.Val == targetSum {
			return true
		}
	}

	if root.Left != nil {
		if left := findLeaf(root.Left, curSum+root.Val, targetSum); left {
			return true
		}
	}
	if root.Right != nil {
		if right := findLeaf(root.Right, curSum+root.Val, targetSum); right {
			return true
		}
	}

	return false
}
