package main

import "fmt"
import "math"

func findComplement(num int) int {
	// 二进制化
	sum := []int{}
	for num > 0 {
		pop := num % 2
		num = num / 2
		sum = append(sum, pop)
	}
	// 逆序
	fmt.Println(sum)
	correct := []int{}
	for i := len(sum) - 1; i >= 0; i-- {
		if sum[i] == 0 {
			sum[i] = 1
		} else {
			sum[i] = 0
		}
		correct = append(correct, sum[i])
	}
	fmt.Println(correct)

	total := 0
	j := 0
	for i := len(correct) - 1; i >= 0; i-- {
		total = total + int(math.Pow(2, float64(j)))*correct[i]
		j++
	}
	fmt.Println(total)

	return total
}

func main() {
	res := findComplement(5)
	fmt.Println(res)
}
