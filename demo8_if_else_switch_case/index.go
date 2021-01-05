package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue > 10 || someValue < 20 {
		fmt.Println("someValue > 10 || someValue < 20")
	} else {
		fmt.Println("NOT someValue > 10 || someValue < 20")
	}

	if result := doSomething(); result == "ok" {
		fmt.Println("OK")
	} else {
		fmt.Println("Nok")
	}

	fnSwitchCase()
}

func doSomething() string {
	return "ok"
}

func fnSwitchCase() {
	i := 3
	switch i {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("bla bla")
		break
	}
}
