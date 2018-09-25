package main

import (
	"fmt"
	"strings"
)

func validPalindrome(s string) bool {
	s = strings.ToLower(s)
	newStr := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			newStr = newStr + string(s[i])
		}
	}
	size := len(newStr) - 1
	flag := false
	for i := 0; i < len(newStr); i++ {
		if size >= 0 && string(newStr[i]) != string(newStr[size]) {
			fmt.Println(string(newStr[i]), string(s[size]))
			if flag {
				return false
			}
			if string(newStr[i]) == string(newStr[size-1]) {

				flag = true
				size--
			} else if i < len(newStr) && string(newStr[i+1]) == string(newStr[size]) {
				//fmt.Println(string(newStr[i+1]), string(s[size]))
				flag = true
				i++
			} else {
				return false
			}
		}
		size--
	}
	return true

}

func main() {
	fmt.Println(validPalindrome("abcaaaa"))
}
