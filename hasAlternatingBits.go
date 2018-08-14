package main

import "fmt"

func hasAlternatingBits(n int) bool {
	arr := []int{}
	for n > 0 {
		pop := n % 2
		n = n / 2
		arr = append(arr, pop)
	}

	sum := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		sum = append(sum, arr[i])
	}

	fmt.Println(len(sum))
	for i := 0; i < len(sum)-1; i++ {
		if sum[i] == sum[i+1] {
			return false
		}
	}
	return true
}

func main() {
	res := hasAlternatingBits(10)
	fmt.Println(res)
}
