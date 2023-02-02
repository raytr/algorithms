package max_even_sum

/*
	we have negativeArr and positiveArr
	iterate all nums and write to 2 arrs
	if sum of all numbers in positive array is even => return this sum
	otherwise we have 2 case:
	- minus a positive number
	- add a negative number


*/

func getMaximumEvenSum(val []int32) int64 {
	eventMax := int32(-1000000000)
	positiveArr := make([]int32, 0, len(val))
	negativeArr := make([]int32, 0, len(val))
	for _, num := range val {
		if num > 0 {
			positiveArr = append(positiveArr, num)
		} else {
			negativeArr = append(negativeArr, num)
		}
	}

	//get sum of number in positive array
	sum := int32(0)
	for _, num := range positiveArr {
		sum += num
	}
	if sum%2 == 0 {
		return int64(sum)
	}

	for _, num := range positiveArr {
		res := sum - num
		if res > eventMax && res%2 == 0 {
			eventMax = res
		}
	}

	for _, num := range negativeArr {
		res := sum + num
		if res%2 == 0 && res > eventMax {
			eventMax = res
		}
	}

	return int64(eventMax)
}
