package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	problem: https://leetcode.com/problems/diameter-of-binary-tree/

    longest path should consist the node, its longest left branch and longest right branch

    call longestPath(u)

    in func longestPath
    node == nil => 0
    left = longestPath(node.left)
    right = longestPath(node.right)
    max = Max(max, left+right)
    return Max(left, right) + 1

*/

var max = 0

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	if root == nil {
		return 0
	}
	longestPathDfs(root)
	return max
}

func longestPathDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := longestPathDfs(root.Left)
	right := longestPathDfs(root.Right)

	max = getMax(max, left+right)
	return getMax(left, right) + 1 //get longest of this branch
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
