package min_cost_climbing_stairs

import "fmt"

/*

   stage:
       at a stair, you have 2 choises, pay the cost or skip
       we have f[i] is sum of cost at step ith
       f[i] = min(f[i-2], f[i-1]) + A[i]

   base case: f[0] = cost[0]
               f[1] = cost[1]
*/

var f []int

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f = make([]int, n+1, n+1)
	cost = append(cost, 0) //represent for top
	f[0], f[1] = cost[0], cost[1]
	r := call(n, cost)
	return r
}

func call(i int, cost []int) int {
	if i == 0 {
		return f[0]
	}
	if i == 1 {
		return f[1]
	}

	if f[i-1] == 0 {
		f[i-1] = call(i-1, cost)
	}
	if f[i-2] == 0 {
		f[i-2] = call(i-2, cost)
	}
	c := f
	fmt.Println(c)
	f[i] = Min(cost[i-1], cost[i-2]) + cost[i]
	a := f[i]
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
