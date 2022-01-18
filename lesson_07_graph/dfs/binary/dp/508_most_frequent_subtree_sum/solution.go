package most_frequent_subtree_sum

//package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	problem: https://leetcode.com/problems/most-frequent-subtree-sum/

	let call sumValueOfNode
	create sumMap
	maxFreq = 0
	result []int

	func sumValOfNode:
	- !node => return 0
	- node.left = sumValOfNode(node.Left)
	- node.right ...
	- sumMap[val+left+right]++

	iterage all of freqMap
	if sumMap[sum] > maxFreq => result = [sum]
	if sumMap[sum] = maxFreq => result.append(sum)
	return result

	complexity:O(n)
*/

var freqMap map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	maxSum := 0
	result := []int{}
	freqMap = make(map[int]int)
	dfs(root)

	for k, v := range freqMap {
		if v > maxSum {
			maxSum = v
			result = []int{k}
		} else if v == maxSum {
			result = append(result, k)
		}
	}

	return result
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	sum := node.Val + left + right
	freqMap[sum]++
	return sum
}
