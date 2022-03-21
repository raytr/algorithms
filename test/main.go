package main

import "fmt"

func main() {
	a := "1234567"
	b := a

	for i, r := range a {
		if i == 2 {
			r = 'N'
			fmt.Println(r)
		}
	}

	fmt.Println(a)
	fmt.Println(b)
}
