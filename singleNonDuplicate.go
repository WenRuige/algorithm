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

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	res := singleNonDuplicate(nums)
	fmt.Println(res)
}
