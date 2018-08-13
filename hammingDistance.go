package main

import (
	"fmt"
	"math"
)

// 汉明距离
func hammingDistance(x int, y int) int {
	num1 := []int{}
	num2 := []int{}
	num3 := []int{}
	num4 := []int{}
	for x > 0 {
		pop := x % 2
		x = x / 2
		num1 = append(num1, pop)
	}

	for y > 0 {
		pop := y % 2
		//fmt.Println(pop)
		y = y / 2
		num2 = append(num2, pop)

	}

	//fmt.Println(len(num1))
	// 逆序

	for i := len(num1) - 1; i >= 0; i-- {
		num3 = append(num3, num1[i])

	}
	num1 = num3
	for i := len(num2) - 1; i >= 0; i-- {
		num4 = append(num4, num2[i])

	}
	num2 = num4

	//fmt.Println(num3, num4)
	//
	//fmt.Println(num1, num2)
	// 补全
	data := len(num2) - len(num1)
	temp := []int{}
	for i := 0; i < int(math.Abs(float64(data))); i++ {
		temp = append(temp, 0)
	}
	if data == 0 {
		// 别操作了
	} else if data > 0 {
		result := make([]int, len(num2))
		res := copy(result, temp)
		copy(result[res:], num1)
		num1 = result
	} else {
		result := make([]int, len(num1))
		res := copy(result, temp)
		copy(result[res:], num2)
		num2 = result

	}

	fmt.Println(num1, num2)
	total := 0
	for i := 0; i < len(num1); i++ {
		if num1[i] != num2[i] {
			total = total + 1
		}
	}
	return total

}

func main() {
	res := hammingDistance(1, 4)
	fmt.Println(res)
}
