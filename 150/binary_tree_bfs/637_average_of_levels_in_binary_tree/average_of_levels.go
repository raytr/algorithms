package average_of_levels_in_binary_tree

/*
	problem : https://leetcode.com/problems/average-of-levels-in-binary-tree
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	res := []float64{float64(root.Val)}

	for len(queue) > 0 {
		sum := 0
		count := 0

		for len(queue) > 0 {
			n := len(queue)
			// pop the node
			cur := queue[n-1]
			queue = queue[:n-1]

			if cur.Left != nil {
				sum += cur.Left.Val
				count++
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				sum += cur.Right.Val
				count++
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, float64(sum)/float64(count))
	}

	return res
}
