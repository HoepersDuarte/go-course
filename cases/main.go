package main

import "fmt"

func main() {
	b := 0

	if b != 0 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	number := 2

	switch number {
	case 2:
		fmt.Println("nunber 2")
	default:
		fmt.Println("defaut")
	}

}
