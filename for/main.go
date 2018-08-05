package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for kay, value := range arr {
		fmt.Println(kay, value*value)
	}

	user := map[string]string{
		"name": "gabriel",
		"nick": "xing",
	}

	for key, value := range user {
		fmt.Printf("o campo \"%s\" tem o valor %s \n", key, value)
	}
}
