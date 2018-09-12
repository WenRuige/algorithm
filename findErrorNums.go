package main

import "fmt"

func findErrorNums(nums []int) []int {
	res := []int{}
	total := make(map[int]int)
	result := 0
	for i := 0; i < len(nums); i++ {
		_, ok := total[nums[i]]
		if ok {
			res = append(res, nums[i])
		} else {
			total[nums[i]] = 1
			// 递增
			result = result + nums[i]
			fmt.Println(i)
		}
	}

	final := ((1 + len(nums)) * len(nums)) / 2
	res = append(res, final-result)
	return res
}

func main() {
	nums := []int{1, 2, 2, 4}
	res := findErrorNums(nums)
	fmt.Println(res)
}
