package convert_sorted_array_to_binary_search_tree

/*

	problem: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
	time complexity: O(n). N is length of nums

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}

	p := (left + right) / 2

	root := &TreeNode{Val: nums[p]}
	root.Left = dfs(left, p-1, nums)
	root.Right = dfs(p+1, right, nums)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(0, len(nums)-1, nums)
}
