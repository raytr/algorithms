package same_tree

/*
	problem: https://leetcode.com/problems/same-tree/

	from the root, i will spread it out to check : is 2 tree are same each layer
	i use breadth first search to solve this problem

	i use a 2 queue to store the index of node being processed in 2 tree

	with each step, i will compare 2 current position of 2 trees
	=> if they not same each other => return false

	when my queue is empty => return true

	runtime: O(n)
	extra space: O(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTreeRecursion(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTreeRecursion(p.Left, q.Left) && isSameTreeRecursion(p.Right, q.Right)
}

func isSameTreeIteration(p *TreeNode, q *TreeNode) bool {

	//edge cases
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	queueP := []*TreeNode{p}
	queueQ := []*TreeNode{q}

	for len(queueP) > 0 && len(queueQ) > 0 {
		curP := queueP[len(queueP)-1]
		curQ := queueQ[len(queueQ)-1]
		if curP.Val != curQ.Val ||
			(curP.Left != nil && curQ.Left == nil) || (curP.Left == nil && curQ.Left != nil) ||
			(curP.Right != nil && curQ.Right == nil) || (curP.Right == nil && curQ.Right != nil) {
			return false
		}

		queueP = queueP[:len(queueP)-1] //remove the last element of queue
		queueQ = queueQ[:len(queueQ)-1] //remove the last element of queue

		if curP.Left != nil {
			queueP = append(queueP, curP.Left)
		}
		if curP.Right != nil {
			queueP = append(queueP, curP.Right)
		}

		if curQ.Left != nil {
			queueQ = append(queueQ, curQ.Left)
		}
		if curQ.Right != nil {
			queueQ = append(queueQ, curQ.Right)
		}
	}

	return true
}
