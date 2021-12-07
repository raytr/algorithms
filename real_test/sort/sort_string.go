package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	sort.Strings(arr)
	fmt.Println(arr)
}
