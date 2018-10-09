package main

import "fmt"

func removeElement(nums []int, val int) int {
	// 快慢指针
	fast, slow := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	nums = nums[0:slow]
	return slow
}

func main() {
	nums := []int{3, 2, 2, 3}
	res := removeElement(nums, 3)
	fmt.Println(res)
}
