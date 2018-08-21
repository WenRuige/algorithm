package main

import "fmt"

func twoSum2(numbers []int, target int) []int {

	var result []int
	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				if i == j {
					continue
				}
				result = []int{i, j}
				for i := 0; i < len(result); i++ {
					result[i] = result[i] + 1
				}
				return result
			}
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum2(nums, 9)
	fmt.Println(res)
}
