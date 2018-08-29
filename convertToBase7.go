package main

import (
	"fmt"
	"math"
	"strconv"
)

func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}
	str := ""
	flag := false
	if num < 0 {
		flag = true
	}
	num = int(math.Abs(float64(num)))
	for num > 0 {
		pop := num % 7
		num = num / 7
		str = str + strconv.Itoa(pop)
	}
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr = newStr + string(str[i])
	}
	if flag {
		newStr = "-" + newStr
	}

	return newStr

}

func main() {
	res := convertToBase7(2)
	fmt.Println(res)
}
