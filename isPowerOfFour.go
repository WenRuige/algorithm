package main

import "fmt"

func isPowerOfFour(num int) bool {

	if num <= 0 {
		return false
	}
	for num > 1 {
		pop := num % 4
		if pop >= 1 {
			return false
		}
		num = num / 4
	}
	return true
}

func main() {
	res := isPowerOfFour(3)
	fmt.Println(res)
}
