package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		pop := n % 3
		if pop != 0 {
			return false
		}
		n = n / 3
	}
	return true
}
func main() {
	flag := isPowerOfThree(-3)
	fmt.Println(flag)
}
