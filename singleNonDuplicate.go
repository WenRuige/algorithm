package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	nu := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := nu[nums[i]]
		if ok {
			// + 1
			nu[nums[i]] = nu[nums[i]] + 1
		} else {
			// 塞里
			nu[nums[i]] = 1
		}
	}

	for k, v := range nu {
		if v == 1 {
			return k
		}
	}

	return 0
}

// 1,1,2,3,3,4,4,8,8
// 有序数组,那么重复的一定是在偶数位置上
func singleNonDuplicate2(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			sum = sum + nums[i]
		} else {
			sum = sum - nums[i]
		}
	}
	return sum
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	res := singleNonDuplicate2(nums)
	fmt.Println(res)
}
