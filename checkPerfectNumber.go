package main

import "fmt"

// time limit
func checkPerfectNumber(num int) bool {
	nums := []int{}
	for i := 1; i < num; i++ {
		if num%i == 0 && i != num {
			nums = append(nums, i)
		}
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
	}
	if total == num {
		return true
	}
	return false
}

func main() {
	res := checkPerfectNumber(99999993)
	fmt.Println(res)
}
