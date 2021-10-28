package binary_tree_level_order_traversal

/*
	problem: https://leetcode.com/problems/binary-tree-level-order-traversal/

	q[1]

	size = 1 =>
	for 1 time {
		3 => level = [1]
		q= [9,20]
	}
	push level to result

	lenght of queue == 2 =>
	for 2 time {
		curNode = q.Front()
		level.push(curNode.Val)
		push(left, right,...) to queue
		q.dequeue()
	}
	result.push(level)


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
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := []int{}
		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = append([]*TreeNode{}, queue[1:]...)
		}
		result = append(result, level)
	}
	return result
}
