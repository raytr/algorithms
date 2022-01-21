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

func levelOrder1(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		level := []int{}
		for _, curNode := range queue {
			level = append(level, curNode.Val)
			queue = append([]*Node{}, queue[1:]...)
			queue = append(queue, curNode.Children...)
		}
		result = append(result, level)
	}
	return result
}

func levelOrder(root *Node) [][]int {
	queue := []*Node{root}
	res := [][]int{}
	if root == nil {
		return res
	}

	for len(queue) > 0 {
		level := []int{}
		for _, curNode := range queue {
			// pop node
			queue := queue[1:]

			level = append(level, curNode.Val)
			// add node to queue
			queue = append(queue, curNode.Children...)
		}
		res = append(res, level)
	}
	return res
}

//q1 := []int{1,2,3,4}
//r1 := []string{}
//
//q2 := []int{1,2,3,4}
//r2 := []string{}
//
//for i:=0; i<len(q1); i++ {
//if len(q1) <10 {
//q1 = append(q1, q1[i]+10)
//}
//r1 = append(r1, fmt.Sprintf("%va",q1[i]))
//}
//
//for _, n := range q2 {
//if len(q1) <10 {
//q1 = append(q1, n+10)
//}
//r2 = append(r2, fmt.Sprintf("%va",n))
//}
//
//fmt.Println(r1)
//fmt.Println(r2)
