package evaluate_boolean_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	//convert leaf Node to boolean
	if root == nil {
		return false
	}
	if root.Val == 0 || root.Val == 1 {
		return root.Val == 1
	}

	//calculate
	leftNode := evaluateTree(root.Left)
	rightNode := evaluateTree(root.Right)

	if root.Val == 2 {
		return leftNode || rightNode
	} else {
		return leftNode && rightNode
	}
}
