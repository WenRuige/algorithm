package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	for n > 0 {
		pop := n % 10
		n = n / 10

		newNum := math.Pow(float64(pop), 2) + math.Pow(float64(n), 2)
		fmt.Println(newNum)
		//os.Exit(1)
	}
	return false
}

func main() {
	res := isHappy(19)
	fmt.Println(res)
}
