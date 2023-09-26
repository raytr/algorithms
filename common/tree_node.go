package common

import (
	"errors"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(arr []*int, i int) *TreeNode {

	if i < len(arr) && arr[i] != nil {
		node := &TreeNode{Val: *arr[i]}
		node.Left = constructTree(arr, 2*i+1)
		//if node.Left == nil {
		//	node.Right = constructTree(arr, 2*i+1)
		//} else {
		node.Right = constructTree(arr, 2*i+2)
		//}
		return node
	}
	return nil
}

func IntegerStringToTreeNode(arrStr string) (*TreeNode, error) {
	if arrStr == "" || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		return nil, errors.New("wrong format")
	}
	arrStr = arrStr[1 : len(arrStr)-1]

	arr := make([]*int, 0)

	nums := strings.Split(arrStr, ",")
	for _, nStr := range nums {
		if nStr == "" || nStr == "null" || nStr == "nil" {
			arr = append(arr, nil)
		} else {
			num, err := strconv.Atoi(nStr)
			if err != nil {
				return nil, err
			}
			arr = append(arr, &num)
		}
	}

	result := constructTree(arr, 0)
	return result, nil

}
