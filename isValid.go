package main

import (
	"fmt"
)

/**
有效的问号?

*/
//func isValid(s string) bool {
//	// 如果不是偶数,肯定无法闭合
//	if len(s)%2 != 0 {
//		return false
//	}
//
//	if len(s) == 2 {
//		if string(s[0]) == "(" {
//			if string(s[1]) != ")" {
//				return false
//			}
//		}
//		if string(s[0]) == "[" {
//			if string(s[1]) != "]" {
//				return false
//			}
//		}
//		if string(s[0]) == "{" {
//			if string(s[1]) != "}" {
//				return false
//			}
//		}
//
//	}
//	// 这个相当于左半部分校验
//	for i := 0; i < len(s)/2; i++ {
//		if i == 0 {
//			if string(s[i]) == "}" || string(s[i]) == "]" || string(s[i]) == ")" {
//				return false
//			}
//
//		} else {
//			if string(s[i]) == "}" && string(s[i-1]) != "{" || string(s[i]) == ")" && string(s[i-1]) != "(" || string(s[i]) == "]" && string(s[i-1]) != "[" {
//				return false
//			}
//		}
//
//	}
//
//	for i := 0; i < len(s)/2; i++ {
//
//		if string(s[i]) == "(" && string(s[i+2]) == ")" {
//			return false
//		}
//
//	}
//	// 校验右半部分
//
//	return true
//}

//type stack []string

func isValid(s string) bool {

	if s == "" {
		return true
	}
	if len(s) < 2 {
		return false
	}
	var stack []string
	for _, v := range s {

		switch string(v) {
		case "}":
			if len(stack) <= 0 {
				return false
			}
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if res != "{" {
				return false
			}
			break
		case "]":
			if len(stack) <= 0 {
				return false
			}
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if res != "[" {
				return false
			}
			break
		case ")":

			if len(stack) <= 0 {
				return false
			}
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if res != "(" {
				return false
			}
			break
		default:
			//fmt.Println(string(v))
			stack = append(stack, string(v))
			break
		}

	}
	if len(stack) > 0 {

		return false
	}

	return true
}

func Push() {

}

func Pop() {

}

func main() {
	//{[]}
	//(]
	//([)]
	res := isValid("()")
	fmt.Println(res)
}
