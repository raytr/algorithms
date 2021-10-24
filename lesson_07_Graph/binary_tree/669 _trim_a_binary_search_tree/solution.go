package trim_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
  with each node

	if node.Val == nil => return nil
	if node.Val > high => return left (call again this func) trimBST(node.Left)
	if node.Val < low => return right (call again this func) trimBST(node.Right)

	with left => check left => left = trimBST(node.Left)
	with right => check right => right = trimBST(node.Right)
	return current node

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
