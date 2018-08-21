package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {

	if n == 0 || n < 0 {
		return false
	}
	for n > 2 {
		pop := n % 2

		if pop != 0 {
			return false
		}
		n = n / 2
	}

	return true

}

func main() {
	res := isPowerOfTwo(16)
	fmt.Println(res)
}
