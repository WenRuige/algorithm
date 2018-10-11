package main

import "fmt"

// 最长回文串

func longestPalindrome(s string) int {

	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
	return 0

}

func main() {
	res := longestPalindrome("abccccdd")
	fmt.Println(res)
}
