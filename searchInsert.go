package main

import "fmt"

func searchInsert(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return len(nums)
}

func main() {

	nums := []int{1, 3, 5, 6}
	res := searchInsert(nums, 5)
	fmt.Println(res)
}
