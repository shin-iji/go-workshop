package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf("Y = %d", y)
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
	fmt.Println(nextInts())

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	seqString := stringSeq()
	fmt.Println(seqString())
	fmt.Println(seqString())
	fmt.Println(stringSeq()())
}
