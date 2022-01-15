package path_sum_ii

/*

	problem: https://leetcode.com/problems/path-sum-ii/

   use dfs, run from root
   with each dfs, we store nodes travelled => meet leaf sum it,
   if == target => append to res

   O(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	res = make([][]int, 0)
	dfs(root, []int{}, 0, targetSum)

	return res
}

func dfs(curNode *TreeNode, curPath []int, sum, targetSum int) {
	if curNode == nil {
		return
	}

	path := append([]int{}, curPath...)
	path = append(path, curNode.Val)
	sum += curNode.Val

	if curNode.Left == nil && curNode.Right == nil {
		if sum == targetSum {
			res = append(res, path)
		}
	}

	dfs(curNode.Left, path, sum, targetSum)
	dfs(curNode.Right, path, sum, targetSum)
}
