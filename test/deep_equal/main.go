package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 4, 2, 3, 1}

	arr3 := []string{"apple", "sammsung", "xiaomi"}
	arr4 := []string{"apple", "xiaomi", "sammsung"}
	arr5 := []string{"apple", "xiaomi", "sammsung"}

	fmt.Println(reflect.DeepEqual(arr1, arr2))
	fmt.Println(reflect.DeepEqual(arr3, arr4))
	fmt.Println(reflect.DeepEqual(arr5, arr4))
	fmt.Println(reflect.DeepEqual(arr1, arr4))
}
