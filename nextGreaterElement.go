package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	result := []int{}
	for i := 0; i < len(findNums); i++ {
		flag := false
		for j := 0; j < len(nums)-1; j++ {
			if findNums[i] == nums[j] {
				for k := j; k < len(nums); k++ {
					if nums[k] > findNums[i] {
						result = append(result, nums[k])
						flag = true
						break
					}
				}

			}
		}
		if !flag {
			result = append(result, -1)
		}
		flag = false
	}
	return result
}

func main() {

	nums1 := []int{3, 1, 5, 7, 9, 2, 6}
	nums2 := []int{1, 2, 3, 5, 6, 7, 9, 11}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println(res)
}
