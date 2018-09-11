package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	first, second, third := -1<<32, -1<<32, -1<<32

	for i := 0; i < len(nums); i++ {
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] < first && nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third && nums[i] < second {
			third = nums[i]
		}
	}
	//if len(nums) >= 3 {
	//	return third
	//} else {
	//	return first
	//}

	fmt.Println(first, second, third)

	if third == -1<<32 || third == second {
		return first
	}
	return third
}

func main() {
	nums := []int{1, 2, 2, 5, 3, 5}
	res := thirdMax(nums)
	fmt.Println(res)
}
