package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnWhileUsingBreak()
}

func fnFor() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("For Index %d\n", i)
	}
}

func fnWhile() {
	i := 0
	for i < 5 {
		fmt.Printf("While Index %d\n", i)
		i++
	}
}

func fnWhileUsingBreak() {
	i := 0
	for true {
		if i > 5 {
			break
		}
		fmt.Printf("WhileBreak Index %d\n", i)
		i++
	}
}
