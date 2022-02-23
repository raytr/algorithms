package main

type MyStruct struct {
}

func a() {
	exMap := make(map[string]int)
	if _, ok := exMap["key"]; ok {
		//do something
	} else {
		// do something
	}
}
