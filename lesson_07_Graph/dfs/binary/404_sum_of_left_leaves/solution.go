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

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	sum = dfs(root, false, sum)
	return sum
}

func dfs(node *TreeNode, isLeft bool, sum int) int {
	if node.Left == nil && node.Right == nil && isLeft {
		return sum + node.Val
	}
	if node.Left != nil {
		sum = dfs(node.Left, true, sum)
	}
	if node.Right != nil {
		sum = dfs(node.Right, false, sum)
	}
	return sum
}
