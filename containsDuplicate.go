package main

import "fmt"

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 1, 2, 3, 4, 5, 6}
	res := containsDuplicate(nums)
	fmt.Println(res)
}
