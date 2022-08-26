package binary_tree_level_order_traversal

/*
	problem: https://leetcode.com/problems/binary-tree-level-order-traversal/

	what is the range of number of nodes in this binary tree
	what is the range of nodes value in this tree

	use bfs to solve it

	queue and push root to queue
	interate all node in queue
		pop node in queue
		append this node to level
		append nodes left and right to queue
		add level to result


	when q.isEmpty => stop while => return result
	complexity is O(m+n) => O(n)
	m is number of node
	n is number of edge

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	for len(queue) > 0 {
		level := make([]int, 0)
		for _, node := range queue {
			queue = queue[1:] //pop node
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
