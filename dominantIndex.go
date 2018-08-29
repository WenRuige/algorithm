package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	// 求一个最大的值
	max := -1 << 32
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			flag = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if max != nums[i] && nums[i] != 0 && max/nums[i] < 2 {
			return -1
		}
	}

	return flag
}

func main() {
	nums := []int{3, 6, 1, 0}
	res := dominantIndex(nums)
	fmt.Println(res)
}
