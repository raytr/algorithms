package main

import "fmt"

//problem: https://leetcode.com/problems/moving-average-from-data-stream/

type MovingAverage struct {
	size int
	arr  []int
}

func main() {
	movingAverage := Constructor(3)
	fmt.Println(movingAverage.Next(1))  // return 1.0 = 1 / 1
	fmt.Println(movingAverage.Next(10)) // return 5.5 = (1 + 10) / 2
	fmt.Println(movingAverage.Next(3))  // return 4.66667 = (1 + 10 + 3) / 3
	fmt.Println(movingAverage.Next(5))  // return 6.0 = (10 + 3 + 5) / 3
}

// complexity: O(size)

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		arr:  []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {

	this.arr = append(this.arr, val)
	if len(this.arr) <= this.size {
		sum := 0
		for _, n := range this.arr {
			sum += n
		}
		return float64(sum) / float64(len(this.arr))
	}

	sum := 0
	for i := len(this.arr) - 1; i >= len(this.arr)-this.size; i-- {
		sum += this.arr[i]
	}
	return float64(sum) / float64(this.size)
}
