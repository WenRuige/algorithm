package main

import (
	"math"
)

func findMin(nums []int) int {
	min := float64(1 << 32)
	for _, v := range nums {
		min = math.Min(min, float64(v))
	}

	return int(min)
}

func main() {
	findMin([]int{5, 6, 7, 1, 2, 3, 4})
}
