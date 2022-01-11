package diameter_of_binary_tree

/*
   define max

   main
       with root == nil => return 0
       dfs(root)
       return max

   use dfs func return number of layer can go down
       left = dfs(node.Left)
       right = dfs(node.right)
       max = getMax(max, left + right)
       return getMax(left, right) + 1 // return deepest node can go down

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max = 0

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs(root)
	return max
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	max = getMax(max, left+right)
	return getMax(left, right) + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
