package main

import "fmt"

func findDuplicates(nums []int) []int {
	total := make(map[int]int)
	newArr := []int{}
	for i := 0; i < len(nums); i++ {
		_, ok := total[nums[i]]
		if ok {
			total[nums[i]] = total[nums[i]] + 1

			if total[nums[i]] == 2 {
				newArr = append(newArr, nums[i])
			}
		} else {
			total[nums[i]] = 1
		}
	}
	return newArr

}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := findDuplicates(nums)
	fmt.Println(res)
}
