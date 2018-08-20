package main

import "fmt"

func missingNumber(nums []int) int {
	if len(nums) == 1 && nums[0] == 1 {
		return 0
	}
	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)

}

func main() {
	nums := []int{0, 1}
	res := missingNumber(nums)
	fmt.Println(res)
}
