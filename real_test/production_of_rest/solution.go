package main

import "fmt"

/*
	given and A & N integer.For each position, you have to print the number of
	pi = product of all elements except A[i]
	A = [2,1,3,4,5]
	Output = 60, 120, 40, 30, 24
-----------------------------------------

	func PrintProd(nums []int)
		prod = product all of number
		for _, num := range A
			print(pro / num)

	time complexity O(n)
	space complexity: O(1)
*/

func main() {
	PrintProd([]int{2, 1, 3, 4, 5})
}

func PrintProd(nums []int) {
	prod := 1
	for _, num := range nums {
		prod *= num
	}

	for _, num := range nums {
		fmt.Println(prod / num)
	}
}
