package capacity_to_ship_packages_within_d_days

/*

problem: https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/


*/

func shipWithinDays(weights []int, days int) int {
	sum, res := 0, 0
	for _, w := range weights {
		sum += w
	}
	lo, hi := 0, sum

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if okWithCapacity(weights, mid, days) {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return res
}

func okWithCapacity(weights []int, capacity, targetDays int) bool {
	neededDays, weightOnShip := 1, 0
	for _, w := range weights {
		if w > capacity {
			return false
		} else if weightOnShip+w <= capacity {
			weightOnShip += w
		} else {
			weightOnShip = w
			neededDays++
		}
	}
	if neededDays > targetDays {
		return false
	}
	return true
}
