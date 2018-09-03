package main

import (
	"fmt"
	"strings"
)

//   only !?',;.
func mostCommonWord(paragraph string, banned []string) string {

	semi := []string{"!", "?", "'", ";", ".", ","}
	paragraph = strings.ToLower(paragraph)
	for i := 0; i < len(semi); i++ {
		flag := strings.Contains(paragraph, semi[i])
		if flag {
			//paragraph = strings.Trim(paragraph, semi[i])
			paragraph = strings.Replace(paragraph, semi[i], "", -1)
		}
	}

	fmt.Println(paragraph)
	str := strings.Split(paragraph, " ")
	strMap := make(map[string]int)
	flag := false
	for i := 0; i < len(str); i++ {
		// 处理
		for j := 0; j < len(banned); j++ {
			if str[i] == banned[j] {
				flag = true
				break
			}
		}
		if !flag {
			_, ok := strMap[string(str[i])]
			if ok {
				strMap[string(str[i])] = strMap[string(str[i])] + 1
			} else {
				strMap[string(str[i])] = 1
			}
		}
		flag = false

	}
	max := 0
	res := ""
	for k, v := range strMap {
		if v > max {
			max = v
			res = k
		}
	}

	fmt.Println(strMap)
	return res
}

func main() {
	str := []string{"hit"}
	res := mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", str)
	fmt.Println(res)
}
