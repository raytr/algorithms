package sum_of_left_leaves

/*
	problem: https://leetcode.com/problems/sum-of-left-leaves/

	create global variable sum = 0

	define dfs:
	- left == nil & right == nil & isLeft => sum += val
	- left != nil => dfs(left, isLeft = true) =>
	- right != nil => dfs(right, isLeft = false)

	complexity: O(2n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	dfs(root, false)
	return sum
}

func dfs(node *TreeNode, isLeft bool) {
	if node.Left == nil && node.Right == nil && isLeft {
		sum += node.Val
	}
	if node.Left != nil {
		dfs(node.Left, true)
	}
	if node.Right != nil {
		dfs(node.Right, false)
	}
}
