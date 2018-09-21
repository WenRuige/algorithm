package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {

	sort.Ints(nums)
	newNums := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			newNums = append(newNums, nums[i])
		}
	}

	fmt.Println(newNums)

	for i := 0; i < len(newNums); i++ {
		if newNums[i] != i+1 {
			fmt.Println(newNums[i], i)
			return i + 1
		}
	}
	return len(newNums) + 1
}

func main() {
	nums := []int{0, 1, 1, 2, 2}
	fmt.Println(firstMissingPositive(nums))
}
