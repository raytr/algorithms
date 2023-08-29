package main

import "fmt"

/*
	we have an array, and a number k. Return list of array that contain all number that sum of them
	less than or equal k
*/

func main() {
	fmt.Println(solution([]int{2, 3, 1, 3}, 6))
}

func solution(arr []int, k int) [][]int {
	allArr := make([][]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i]
		curArr := []int{arr[i]}
		allArr = append(allArr, curArr)
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			curArr = append(curArr, arr[j])
			if sum <= k {
				allArr = append(allArr, curArr)
			}
		}
	}

	fmt.Println(allArr)

	return allArr
}
