package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	value := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := value[nums[i]]
		if ok {
			value[nums[i]] = value[nums[i]] + 1

		} else {
			value[nums[i]] = 1
		}
	}

	for k, v := range value {
		fmt.Println(k, v)
	}

	return []int{}

}

func main() {

	nums := []int{1, 1, 1, 2, 2, 3}
	topKFrequent(nums, 2)
}
