package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	// 二分查找
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(nums[mid])
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	//nums := []int{-1, 0, 3, 5, 9, 12}

	nums := []int{5}
	res := search(nums, 5)
	fmt.Println(res)
}
