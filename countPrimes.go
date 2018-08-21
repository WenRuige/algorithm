package main

import "fmt"

func countPrimes(n int) int {

	if n == 0 || n == 1 || n == 2 {
		return 0
	}
	total := 0
	for i := 2; i <= n; i++ {
		// 除了1 或者本身
		flag := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}

		if !flag {
			total++
		}
		flag = false

	}

	return total

}

func main() {
	res := countPrimes(10)
	fmt.Println(res)
}
