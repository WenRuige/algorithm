package main

import "fmt"

func mySqrt(x int) int {

	if x == 1 || x == 2 || x == 3 {
		return 1
	}

	for i := 0; i < x; i++ {

		if i*i == x {
			return i
		}
		if i*i < x {

		} else {

			return i - 1
		}
	}
	return 0
}

func main() {
	res := mySqrt(17)
	fmt.Printf("ssss   %d", res)
}
