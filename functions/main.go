package main

import (
	"fmt"
)

func main() {
	a, _ := tryFunc()
	_, b := tryFunc()
	fmt.Println(a, b)
}

func tryFunc() (string, string) {
	str1 := "string 1"
	str2 := "string 2"
	return str1, str2
}
