package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	str := []string{}
	temp := ""
	for i := 0; i < len(S); i++ {
		if S[i] != ' ' {
			temp = temp + string(S[i])
		} else {
			str = append(str, temp)
			temp = ""
		}
		if i == len(S)-1 {
			str = append(str, temp)
		}
	}
	// 原因列表
	//newStr := []string{}
	yuanyin := make(map[string]int)
	yuanyin["a"] = 1
	yuanyin["e"] = 1
	yuanyin["i"] = 1
	yuanyin["o"] = 1
	yuanyin["u"] = 1

	yuanyin["A"] = 1
	yuanyin["E"] = 1
	yuanyin["I"] = 1
	yuanyin["O"] = 1
	yuanyin["U"] = 1
	for i := 0; i < len(str); i++ {
		_, ok := yuanyin[string(str[i][0])]
		// 如果是元音字母的话
		if ok {
			str[i] = str[i] + "ma"
		} else {
			str[i] = str[i][1:] + str[i][:1] + "ma"
			// 辅音字母
		}
		for j := 0; j <= i; j++ {
			str[i] = str[i] + "a"
		}
	}
	newStr := ""
	for i := 0; i < len(str); i++ {
		newStr = newStr + str[i] + " "
	}
	newStr = strings.TrimRight(newStr, " ")
	return newStr
}

func main() {
	res := toGoatLatin("I speak Goat Latin")
	fmt.Println(res)
}
