package main

import "fmt"

func maxProfitI(prices []int) int {

	max:=0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if prices[j] < prices[i] {
				continue
			}else{
				if prices[j] - prices[i] > max{
					max = prices[j] - prices[i]
				}
			}
		}
	}
	return max
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfitI(nums))
}
