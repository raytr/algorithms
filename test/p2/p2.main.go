package main

import "fmt"

func main() {
	s1 := 7
	//s1 := 11
	count := 0

	for s1 < 620 {
		sum1 := s1
		for s1 > 1 {
			surplus := s1 % 10
			sum1 += surplus
			s1 /= 10
		}
		s1 = sum1
		count++
	}

	fmt.Println(count)
}
