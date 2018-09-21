package main

import (
	"fmt"
	"strings"
)

//func isPalindrome(x int) bool {
//
//	origin := x
//
//	if x < 0 {
//		return false
//	}
//	res := 0
//	for x > 0 {
//		pop := x % 10
//		x = x / 10
//
//		res = res*10 + pop
//
//	}
//	if res == origin {
//		return true
//	}
//	return false
//}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	newStr := ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] <= '9' && s[i] >= '0') {
			newStr = newStr + string(s[i])
		}
	}
	size := len(newStr) - 1
	for i := 0; i < len(newStr); i++ {
		fmt.Println(string(newStr[i]), string(s[size]))
		if string(newStr[i]) != string(newStr[size]) {
			return false
		}
		size--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
