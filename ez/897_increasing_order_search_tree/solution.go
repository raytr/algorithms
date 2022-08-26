package increasing_order_search_tree

/*
	problem: https://leetcode.com/problems/increasing-order-search-tree/

   clarify question:
       what is the maximum of nodes in this binary tree
       is has duplicate nodes in this tree
       what happen when no node in this tree


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	inorder(root, &vals)

	ans := &TreeNode{
		Val: 0,
	}
	cur := ans
	for _, v := range vals {
		cur.Right = &TreeNode{Val: v}
		cur = cur.Right
	}
	return ans.Right
}

func inorder(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, vals)
	*vals = append(*vals, node.Val)
	inorder(node.Right, vals)
}
