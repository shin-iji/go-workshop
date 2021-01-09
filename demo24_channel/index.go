package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("step1")
	fmt.Println(<-ch)

	ch <- 2
	fmt.Println("step2")
	fmt.Println(<-ch)

}
