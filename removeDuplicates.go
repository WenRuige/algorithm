package main

import "fmt"

func removeDuplicates(nums []int) int {
	number := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[number] {
			number++
			nums[number] = nums[i]
		}
	}
	return number + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
