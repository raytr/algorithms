package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(arr []*int, i int) *TreeNode {
	if i == 2 {
		fmt.Println(true)
	}

	if i < len(arr) && arr[i] != nil {
		node := &TreeNode{Val: *arr[i]}
		i1 := 2*i + 1
		node.Left = constructTree(arr, i1)
		if node.Left == nil && i > 0 {
			i2 := 2 * i
			node.Right = constructTree(arr, i2)
		} else {
			i2 := 2*i + 2
			node.Right = constructTree(arr, i2)
		}
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
			//arr = append(arr, []*int{nil, nil, nil}...) //root, left, right
			arr = append(arr, nil) //root, left, right
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

func main() {
	result, _ := IntegerStringToTreeNode("[1,null,2,3]")
	fmt.Println(fmt.Sprintf("%v", result))
}
