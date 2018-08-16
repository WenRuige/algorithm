package main

import "fmt"

func titleToNumber(s string) int {
	res := 0
	tmp := 1
	for i := len(s) - 1; i >= 0; i-- {
		res = res + int((s[i]-'A'+1))*tmp
		tmp = tmp * 26
	}
	fmt.Println(res)
	return res
}

func main() {
	res := titleToNumber("AB")
	fmt.Println(res)
}
