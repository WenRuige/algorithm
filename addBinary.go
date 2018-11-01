package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {

	if len(a) > len(b) {
		res := len(a) - len(b)
		for i := 0; i < res; i++ {
			b = "0" + b
		}
	}
	if len(a) < len(b) {
		res := len(b) - len(a)
		for i := 0; i < res; i++ {
			a = "0" + a
		}
	}
	// 前置补0
	str := ""
	flag := 0
	for i := len(a) - 1; i >= 0; i++ {
		int1, _ := strconv.Atoi(string(a[i]))
		int2, _ := strconv.Atoi(string(b[i]))
		if int1+int2 > 2 {
			// 二进制与操作进0
			flag = 1
			str = str + "0"
		} else {
			if flag == 1 {
				str = str + "0"
			} else {
				str = str + "1"
			}

		}
	}

	fmt.Println(a, b)
	return ""
}

func main() {
	res := addBinary("1011", "1011")
	fmt.Println(res)
}
