package main

import "fmt"

// 最长回文串

func longestPalindrome(s string) int {
	str := make(map[string]int)
	for i := 0; i < len(s); i++ {
		_, ok := str[string(s[i])]
		if ok {
			str[string(s[i])] = str[string(s[i])] + 1
		} else {
			str[string(s[i])] = 1
		}

	}
	final := 0
	flag := 0
	for _, v := range str {
		//fmt.Println(k, v)
		if v%2 == 0 {
			final = final + v
		} else {
			final = final + v - 1
			flag = 1
		}
	}

	if flag == 1 {
		return final + 1
	}

	return final

}

func main() {
	res := longestPalindrome("ccc")
	fmt.Println(res)
}
