package remove_element

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	reader := 0
	writer := 0
	k := n

	for reader < n {

		if nums[reader] == val || nums[writer] == val {
			k--
			//move reader until meet num != val
			for reader < n && nums[reader] == val {
				reader++
			}

			if reader < n {
				nums[writer] = nums[reader]
			}
		}
		reader++
		writer++
	}

	fmt.Println(nums)
	return k
}
