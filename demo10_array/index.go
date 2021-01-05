package main

import "fmt"

func main() {
	var array1 = []int{1, 2, 3, 4}
	var array2 [3]string

	for _, item := range array1 {
		fmt.Print(item, ",")
	}

	array2[0] = "Android"
	array2[1] = "iOS"
	array2[2] = "React"
	for _, item := range array2 {
		fmt.Print(item, ",")
	}
}
