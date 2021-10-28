//package k_closest_points_to_origin
package main

import (
	"fmt"
	"github.com/raytr/algorithms/lesson_03/quick_select/k_smallest_element_in_an_array"
)

func main() {
	//[[6,10],[-3,3],[-2,5],[0,2]]
	//fmt.Println(kClosest([][]int{{6, 10}, {-3, 3}, {-2, 5}, {0, 2}}, 3))
	//[[]]
	//	fmt.Println(kClosest([][]int{{-2,10},{-4, -8},{10, 7},{-4,-7}}, 3))
	//fmt.Println(kClosest([][]int{{10, -2}, {2, -2}, {10, 10}, {9, 4}, {-8, 1}}, 4))
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}

//problem: https://leetcode.com/problems/k-closest-points-to-origin/
func kClosest(points [][]int, k int) [][]int {
	// first change points array to distance square array and store in a map
	pointMap := make(map[int][][]int)
	distanceSqrArr := make([]int, 0, len(points))
	for _, p := range points {
		d2 := p[0]*p[0] + p[1]*p[1]

		if _, exist := pointMap[d2]; !exist {
			pointMap[d2] = [][]int{}
		}
		pointMap[d2] = append(pointMap[d2], p)

		distanceSqrArr = append(distanceSqrArr, d2)
	}

	//use quick select to get 2 smallest distanceSqr
	distanceSqrArr = k_smallest_element_in_an_array.FindKNumsSmallest(distanceSqrArr, k)

	// store in points
	points = points[:0]
	for _, d2 := range distanceSqrArr {
		points = append(points, pointMap[d2]...)
	}

	return points[:k]
}
