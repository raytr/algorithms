package delete_node_in_a_bst

//problem: https://leetcode.com/problems/delete-node-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var firstLeaf *TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	/*
	   1. if target has no child, simply remote that node
	   2. if target has 1 child, use its child to replace that node
	   3. if target has 2 child, replace that node with its in-order successor and remove that node
	*/

	if root == nil || root.Val == key {
		return delete(root)
	}
	if key > root.Val {
		return deleteNode(root.Right, key)
	} else {
		return deleteNode(root.Left, key)
	}
}

func delete(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		root = nil
	} else if root.Left != nil && root.Right == nil {
		root = root.Left
	} else if root.Right != nil && root.Left == nil {
		root = root.Right
	} else {
		findLeafDfs(root)
		root.Val = firstLeaf.Val
		firstLeaf = nil
	}

}

func findLeafDfs(root *TreeNode, leftLeaf, rightLeaf bool) *TreeNode {
	if root.Left != nil && root.Right != nil {
		if leftLeaf {

		}
	}
	if root.Left != nil {
		findLeafDfs(root.Left)
	}
	if root.Right != nil {
		findLeafDfs(root.Right)
	}

}
