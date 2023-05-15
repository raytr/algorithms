package n_ary_tree_level_order_traversal

/*
	problem: https://leetcode.com/problems/n-ary-tree-level-order-traversal/

	q = []node
	result = [][]int
	bfs =>
	loop !q.isEmpty{
		level is []int
		for {
			curNode = q[0]
			level.push(curNode)
			dequeue
			queue.add(node.Children)
		}
		result.push(level)
	}

	return result



*/

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := []*Node{root}
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	for len(queue) > 0 {
		level := make([]int, 0)
		for _, node := range queue {
			queue = queue[1:] //pop node
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		res = append(res, level)
	}
	return res
}
