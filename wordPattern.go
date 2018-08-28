package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	// 首先将str 拆分

	res := strings.Split(str, " ")
	//if len(res) == 1 {
	//	if pattern == str {
	//		return true
	//	}
	//}
	if len(pattern) != len(res) {
		return false
	}
	total := make(map[string]string)

	for i := 0; i < len(pattern); i++ {
		_, ok := total[string(pattern[i])]
		if ok {
			if total[string(pattern[i])] != res[i] {
				return false
			}

		} else {
			total[string(pattern[i])] = res[i]
			for k, v := range total {
				//fmt.Println(k, v)
				if v == res[i] && k != string(pattern[i]) {
					return false
				}
			}

		}
	}
	return true

}

func main() {
	res := wordPattern("jq", "jq")
	fmt.Println(res)
}
