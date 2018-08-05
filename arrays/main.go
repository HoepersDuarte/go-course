package main

import (
	"fmt"
)

func main() {
	tabuada := [10]int{0, 5, 10}

	tabuada[3] = 15

	fmt.Println(tabuada)

	user := map[string]string{
		"name": "eu",
		"nick": "xing",
	}

	user["age"] = "20"

	fmt.Println(user)

	slice := []int{10, 10}
	slice = append(slice, 90, 110, 50)

	fmt.Println(slice)

	sl := make([]int, 2, 3)
	sl[0] = 90
	sl[1] = 10

	sl2 := sl[:]

	sl2[0] = 1000

	fmt.Println(sl)
	fmt.Println(sl2)

}
