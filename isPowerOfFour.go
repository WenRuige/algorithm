package main

import "fmt"

func isPowerOfFour(num int) bool {

	if num == 5 || num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	for num > 1 {
		pop := num % 4
		if pop > 1 && num == 4 {
			return true
		} else if pop == 0 {
			return true
		} else {
			return false
		}
		num = num / 4
	}
	return true
}

func main() {
	res := isPowerOfFour(8)
	fmt.Println(res)
}
