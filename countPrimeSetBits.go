package main

import "fmt"

func countPrimeSetBits(L int, R int) int {
	total := 0

	for i := L; i <= R; i++ {
		nums := []int{}
		num := i
		for num > 0 {
			pop := num % 2
			nums = append(nums, pop)
			num = num / 2
		}
		//质数
		primeNumber := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == 1 {
				primeNumber++
			}
		}
		temp := 0
		for k := 2; k < primeNumber; k++ {
			//fmt.Println(primeNumber, k)
			if primeNumber%k == 0 {
				temp = temp + 1
			}
		}
		if temp == 0 && primeNumber != 1 {
			total = total + 1
		}
	}

	return total
}

func main() {
	res := countPrimeSetBits(244, 269)
	fmt.Println(res)
}
