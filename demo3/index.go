package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("Test")

	var tmp1 int = 0
	var tmp2 string = "Hello"
	var tmp3 bool = true
	tmp4 := 1

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)

	run()
	run()
	run()

}

func run() {
	count++
	fmt.Println(count)
}
