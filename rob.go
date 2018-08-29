package main

import "fmt"

// 动态规划
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	final := -1 << 32
	for i := 0; i < len(nums)-1; i++ {
		total := 0
		for j := i; j < len(nums); j++ {
			total = total + nums[j]
			j = j + 1
		}
		//fmt.Println(total)
		if final < total {
			final = total
		}
	}

	return final

}

func main() {
	nums := []int{2, 1, 1, 2}
	res := rob(nums)
	fmt.Println(res)
}
