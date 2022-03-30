package main

import "fmt"
import "os"

func IsDuoDigit(number int) string {
	// Write your code here
	// fmt.Fprintln(os.Stderr, "Debug messages...")

	if number > 0 {
		m := make(map[int]bool)
		for number > 1 {
			surplus := number % 10
			m[surplus] = true
			number /= 10
		}
		if len(m) > 2 {
			return "n"
		}
		return "y"
	} else {
		m := make(map[int]bool)
		for number < -1 {
			surplus := number % 10
			m[surplus] = true
			number /= 10
		}
		if len(m) > 2 {
			return "n"
		}
		return "y"
	}

	return ""
}

/* Ignore and do not change the code below */
func main() {
	var number int
	fmt.Scan(&number)

	stdoutWriter := os.Stdout
	os.Stdout = os.Stderr
	result := IsDuoDigit(number)
	os.Stdout = stdoutWriter
	fmt.Println(result)
}
