package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {

	for i := 1; num > 0; i = i + 2 {
		num = num - i
	}
	return num == 0
}

func main() {
	res := isPerfectSquare(16)
	fmt.Println(res)
}
