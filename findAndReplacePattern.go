package main

import "fmt"

// 查找和替换模式
func findAndReplacePattern(words []string, pattern string) []string {
	final := []string{}
	for i := 0; i < len(words); i++ {
		if strMatch(words[i], pattern) {
			final = append(final, words[i])
		}
	}
	return final
}

// 俩字符串进行比较
func strMatch(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		for j := i + 1; j < len(str1); j++ {
			if str1[i] == str1[j] && str2[i] != str2[j] {
				return false
			}
			if str1[i] != str1[j] && str2[i] == str2[j] {
				return false
			}
		}

	}
	return true
}

func main() {
	words := []string{"abb", "sss"}
	pattern := "abb"
	res := findAndReplacePattern(words, pattern)
	fmt.Println(res)
}
