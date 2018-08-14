package main

import "fmt"

func addDigits(num int) int {

	if num < 10 {
		return num
	}

	total := 2 << 32
	flag := true
	for total > 10 || flag {
		total = digui(num)
		flag = false
	}

	fmt.Println(total)

	return 0

}

func digui(total int) int {
	other := 0
	for total > 0 {
		pop := total % 10
		total = total / 10
		other = other + pop
	}
	fmt.Println(other)
	return other
}

func main() {
	addDigits(38)
}
