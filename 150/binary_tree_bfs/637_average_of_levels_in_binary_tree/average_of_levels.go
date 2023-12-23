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
	res := make([]float64, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		sum, count := 0, 0

		for len(queue) > 0 {
			//pop a node
			curNode := queue[0]
			queue = queue[1:]
			sum += curNode.Val
			count++

			//push children to temp
			if curNode.Left != nil {
				temp = append(temp, curNode.Left)
			}

			if curNode.Right != nil {
				temp = append(temp, curNode.Right)
			}
		}

		queue = temp
		res = append(res, float64(sum)/float64(count))
	}

	return res
}
