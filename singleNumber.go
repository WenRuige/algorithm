package main

import "fmt"

func singleNumber(nums []int) int {

	if len(nums) <= 1 {
		return nums[0]
	}
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = hashmap[nums[i]] + 1
	}

	fmt.Println(hashmap)
	for v := range hashmap {
		if hashmap[v] == 1 {
			return v

		}
	}
	return 0
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)
	fmt.Println(res)
}
