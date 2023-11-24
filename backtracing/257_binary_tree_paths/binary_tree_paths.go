package binary_tree_paths

import (
	"strconv"
)

/*
	problem: https://leetcode.com/problems/binary-tree-paths
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)
	buildThePaths(root, "", &paths)
	return paths
}

func buildThePaths(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}

	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil { //if root is leaf
		*paths = append(*paths, path)
	} else {
		path += "->"
		buildThePaths(root.Left, path, paths)
		buildThePaths(root.Right, path, paths)
	}
}
