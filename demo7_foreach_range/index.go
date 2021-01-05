package main

import "fmt"

func main() {
	courses := []string{"Android", "iOS", "React"}

	for i, item := range courses {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	for _, item := range courses {
		fmt.Printf("%s\n", item)
	}
}
