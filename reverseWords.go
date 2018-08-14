package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	str := []string{}
	flag := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			temp := s[flag:i]
			flag = i
			temp = strings.Trim(temp, " ")
			str = append(str, temp)
		}

	}
	if flag < len(s) {
		temp := s[flag:]
		strings.Trim(temp, " ")
		str = append(str, temp)
	}
	result := []byte{}
	for i := 0; i < len(str); i++ {
		for j := len(str[i]) - 1; j >= 0; j-- {
			//fmt.Println(len(str[i]))
			fmt.Printf("%v  ,%v \n", i, j)
			result = append(result, str[i][j])
		}
		result = append(result, ' ')
	}

	return strings.TrimRight(string(result), " ")
}

func main() {

	words := `Let's take LeetCode contest`
	res := reverseWords(words)
	fmt.Println(res)
}
