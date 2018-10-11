package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	for i := len(s) / 2; i > 0; i-- {
		// 循环的次数,必须是一个可除开的数
		if len(s)%i == 0 {
			// 子串的数量
			substr := len(s) / i
			// 获取前i个数,看拼起来的是否和本身相同
			sub := s[0:i]
			temp := ""
			for j := 0; j < substr; j++ {
				temp = temp + sub
			}
			if temp == s {
				return true
			}
		}
	}
	return false
}

func main() {
	res := repeatedSubstringPattern("gwrgwrgwrs")
	fmt.Println(res)
}
