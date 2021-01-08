package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Println("Finish", i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
