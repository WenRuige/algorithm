package main

import (
	"fmt"
	"math"
)

// 先找出最大值和最小值
func smallestRangeI(A []int, K int) int {
	max := float64(math.MinInt64)
	min := float64(math.MaxFloat64)
	for i := 0; i < len(A); i++ {
		max = math.Max(max, float64(A[i]))
		min = math.Min(min, float64(A[i]))
	}
	fmt.Println(max, min)
	value := max - float64(K) - (min + float64(K))

	if value >= 0 {
		return int(value)
	} else {
		return 0
	}
}

func main() {

	nums := []int{1, 2, 3}
	fmt.Println(smallestRangeI(nums, 1))
}
