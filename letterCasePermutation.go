package main

import (
	"fmt"
	"strings"
)

func letterCasePermutation(S string) []string {
	result := []string{}

	ddfs("", S, &result, 0)
	fmt.Println(result)
	return result

}

func ddfs(pre, str string, result *[]string, index int) {
	if index == len(str) {
		*result = append(*result, pre)
	} else {
		s := str[index]
		if s <= 122 && s >= 97 || s >= 65 && s <= 90 {
			ch := strings.ToLower(string(str[index]))
			ddfs(pre+ch, str, result, index+1)
			cj := strings.ToUpper(string(str[index]))
			ddfs(pre+cj, str, result, index+1)
		} else {
			ddfs(pre+string(s), str, result, index+1)
		}
	}
}

func main() {

	letterCasePermutation("a1B2")

}
