package main

import "fmt"

func main() {
	s := "Welcome {name} to {company}"
	m := make(map[string]string)
	m["name"] = "Ray Tran"
	m["company"] = "Orbitau"
	fmt.Println(stringReplace(s, m))
}

func stringReplace(s string, m map[string]string) string {
	var temp []int32
	var isTemp bool
	output := ""
	for _, c := range s {
		if c == '{' {
			isTemp = true
			continue
		}
		if c == '}' {
			isTemp = false
			output += m[string(temp)]
			temp = []int32{}
			continue
		}
		if isTemp {
			temp = append(temp, c)
			continue
		}
		output += string(c)
	}
	return output
}
