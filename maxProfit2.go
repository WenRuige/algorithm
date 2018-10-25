package main

import "fmt"

func maxProfit(prices []int) int {
	max := 0
	size := len(prices)
	for i := 0; i < size-1; i++ {
		if prices[i] < prices[i+1] {
			max = max + prices[i+1] - prices[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
}
