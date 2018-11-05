package main

import "fmt"

func nextGreaterElementsV2(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			// 如果顺序找到的话
			if nums[i] < nums[j] {
				result = append(result, nums[j])
				break
			}
		}
		if len(result) == i {
			for j := 0; j < len(nums); j++ {
				// 如果顺序找到的话
				if nums[i] < nums[j] {
					result = append(result, nums[j])
					break
				}
			}
			if len(result) == i {
				result = append(result, -1)
			}

		}

	}

	return result
}

func main() {
	res := nextGreaterElementsV2([]int{1, 3, 2, 1})
	fmt.Println(res)
}
