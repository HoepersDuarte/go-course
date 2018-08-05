package main

import (
	"fmt"
)

func main() {
	number := 20
	for number > 0 {
		fmt.Println(number)
		number--
	}

	for n := 10; n > 0; n-- {
		fmt.Println(n)
	}

}
