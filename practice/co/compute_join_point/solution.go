package compute_join_point

import "fmt"

func main() {
	fmt.Println(ComputeJoinPoint(471, 480))
	fmt.Println(ComputeJoinPoint(11, 7))              //620
	fmt.Println(ComputeJoinPoint(32, 47))             // 47
	fmt.Println(ComputeJoinPoint(1158, 2085))         // 2130
	fmt.Println(ComputeJoinPoint(14689, 13))          // 20014
	fmt.Println(ComputeJoinPoint(991, 997))           // 1118
	fmt.Println(ComputeJoinPoint(15485863, 15215260)) // 1118
}

func ComputeJoinPoint(s1 int, s2 int) int {
	existMap := make(map[int]bool)
	max := 20000000

	for s1 != s2 || s1 < max || s2 < max {
		sum1 := s1
		for s1 >= 1 {
			surplus := s1 % 10
			sum1 += surplus
			s1 /= 10
		}

		sum2 := s2
		for s2 >= 1 {
			surplus := s2 % 10
			sum2 += surplus
			s2 /= 10
		}

		//fmt.Println(fmt.Sprintf("sum1 is: %v", sum1))
		//fmt.Println(fmt.Sprintf("sum2 is: %v", sum2))

		if sum1 == sum2 {
			return sum1
		}
		if _, exist := existMap[sum1]; exist {
			return sum1
		}
		if _, exist := existMap[sum2]; exist {
			return sum2
		}

		s1 = sum1
		s2 = sum2
		existMap[sum1] = true
		existMap[sum2] = true
	}

	return -1
}
