package binary_tree_level_order_traversal_ii

/*

problem: https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

queue []*Node{root}

while !queue.isEmpty {
    level = []int{}
    for n range of queue {
        level.push(n)
        push left => queue (left == !null)
        push right => queue (right == !null)
        dequeue
}
    result.push(level)
}

    return reverse(result)

    O(n)

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := []int{}
		for _, n := range queue {
			level = append(level, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			queue = queue[1:]
		}
		result = append(result, level)
	}
	return reverse(result)
}

func reverse(arr [][]int) [][]int {
	reversed := [][]int{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}
