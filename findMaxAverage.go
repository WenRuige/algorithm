package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	max := float64(-1 << 32)
	if len(nums) == k {
		total := float64(0)
		for i := 0; i < len(nums); i++ {
			total = total + float64(nums[i])
		}
		return total / float64(k)
	}
	for i := 0; i < len(nums) && len(nums)-i >= k; i++ {
		total := float64(0)
		begin := 0
		for j := i; begin < k; begin++ {
			total = total + float64(nums[j])
			j++
		}
		if max < total {
			max = total
		}
	}

	return max / float64(k)
}

func main() {
	nums := []int{0, 1, 1, 3, 3}
	res := findMaxAverage(nums, 4)
	fmt.Println(res)
}
