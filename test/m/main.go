package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	myNode := &Node{
		Val: 1,
	}

	if myNode.Next != nil && myNode.Next.Val != 3 {
		fmt.Println(myNode.Next.Val)
	}

	fmt.Println("finish")

	//if myNode.Next.Val != 3 && myNode.Next != nil {
	//	fmt.Println(myNode.Next.Val)
	//}
	//
	//fmt.Println("finish")

}
