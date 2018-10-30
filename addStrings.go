package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		res := len(num2) - len(num1)
		for i := 0; i < res; i++ {
			num1 = "0" + num1
		}
	} else {
		res := len(num1) - len(num2)
		for j := 0; j < res; j++ {
			num2 = "0" + num2

		}
	}
	fmt.Println(num1, num2)
	newStr := ""
	flag := 0

	for i := len(num1) - 1; i >= 0; i-- {
		int1, _ := strconv.Atoi(string(num1[i]))
		int2, _ := strconv.Atoi(string(num2[i]))
		if int1+int2+flag >= 10 {
			if i == 0  {
				newStr = fmt.Sprint((int1 + int2 + flag)) + newStr
			} else {
				flag ++
				newStr = fmt.Sprint((int1+int2)%10) + newStr
			}
		} else {
			newStr = fmt.Sprint(int1+int2+flag) + newStr
			flag = 0

		}
		//fmt.Println(newStr)
	}
	return newStr
}

func main() {
	res := addStrings("999", "999")
	fmt.Println(res)
}
