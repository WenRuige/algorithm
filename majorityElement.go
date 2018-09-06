package main

import "fmt"

//func majorityElement(nums []int) int {
//	res := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		_, ok := res[nums[i]]
//		if ok {
//			res[nums[i]] = res[nums[i]] + 1
//		} else {
//			res[nums[i]] = 1
//		}
//	}
//
//	flag := 0
//	max := -2 << 32
//	for k, v := range res {
//		if v > max {
//			max = v
//			flag = k
//		}
//	}
//	return flag
//}

func majorityElement(nums []int) []int {
	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := result[nums[i]]
		if ok {
			result[nums[i]] = result[nums[i]] + 1
		} else {
			result[nums[i]] = 1
		}
	}

	flag := len(nums) / 3
	res := []int{}
	for k, v := range result {
		if v > flag {
			res = append(res, k)
		}

	}
	return res

}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	fmt.Println(res)
}
