package main

import (
	"fmt"
	//"os"
	"strings"
)

func lengthOfLastWord(s string) int {

	if len(strings.TrimSpace(s)) == 0 {
		return 0
	}

	s = strings.TrimRight(s, " ")

	fmt.Println(s)

	//if s[len(s)-1] == ' ' {
	//	s = s[0 : len(s)-1]
	//}

	fmt.Println(s)
	// case 1 处理带空格的情况
	for i := len(s) - 1; i >= 0; i-- {
		// 如果空值不是倒数第个的话
		if s[i] == ' ' && (len(s)-1) != i {
			return len(s) - 1 - i
		}

	}
	// case 2:处理没空格的情况
	return len(s)
}

func main() {
	res := lengthOfLastWord("a ")
	fmt.Println(res)
}
