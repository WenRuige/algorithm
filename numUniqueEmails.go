package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	finalStr := make(map[string]int)

	for i := 0; i < len(emails); i++ {
		//实际上只需要左半部分
		str := ""
		newStr := strings.Split(emails[i], "@")

		for j := 0; j < len(newStr[0]); j++ {
			if string(newStr[0][j]) == "." {
				continue
			} else if string(newStr[0][j]) == "+" {
				break

			} else {
				str = str + string(newStr[0][j])
			}

		}
		_, ok := finalStr[str+newStr[1]]
		if !ok {
			finalStr[str+newStr[1]] = 1
		}

	}
	return len(finalStr)
}

func main() {
	res := numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
	fmt.Println(res)
}
