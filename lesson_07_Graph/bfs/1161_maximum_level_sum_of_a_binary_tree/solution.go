package maximum_level_sum_of_a_binary_tree

/*

problem: https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/

root == nil => return 0
queue = []*Node{} => queue.push(root)
maxSum = sum
maxLevelVal = 1
level = 1
maxSum[1] = root.Val

while !queue.IsEmpty {
	sum = 0
	for node := range of queue {
		sum += node.Val
		if left != nul => q.push(left)
		... right
		dequeue
	}
	if sum > maxSum => maxSum = sum; maxLevelVal = level
	level++
}
return maxSum

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 1
	maxLevel := 1
	maxSum := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := 0
		for _, node := range q {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}
		level++
	}
	return maxLevel
}
