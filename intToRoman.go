package main

import (
	"fmt"
)

func intToRoman(num int) string {
	number := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	flag := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	str := ""
	for i := 0; i < 13 && num > 0; i++ {
		if num < number[i] {
			continue
		}

		for num >= number[i] {
			num = num - number[i]
			str = str + flag[i]
		}

	}
	fmt.Println(str)
	return str
}

func main() {
	res := intToRoman(230)
	fmt.Println(res)
}
