package main

import "fmt"
import "os"

func Solve(width int, height int, length int, mass int) string {
	// Write your code here
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(width, height, length, mass)

	return ""
}

/* Ignore and do not change the code below */
func main() {
	for {
		var width, height, length, mass int
		fmt.Scan(&width, &height, &length, &mass)

		stdoutWriter := os.Stdout
		os.Stdout = os.Stderr
		action := Solve(width, height, length, mass)
		os.Stdout = stdoutWriter
		fmt.Println(action)
	}
}

//160 70 70 10
//90 90 90 10
//120 120 120 30
//150 70 70 10
