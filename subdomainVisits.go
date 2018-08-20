package main

import (
	"fmt"
	//"os"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {

	str := make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		res := strings.Split(cpdomains[i], " ")
		data, _ := strconv.Atoi(res[0])

		temp := strings.Split(res[1], ".")

		flag := 0
		for i := len(temp) - 1; i >= 0; i-- {
			newStr := ""
			for j := flag; j < len(temp); j++ {
				newStr = newStr + string(temp[j]) + "."
			}
			newStrTrim := strings.TrimRight(newStr, ".")
			_, ok := str[newStrTrim]
			if ok {
				str[newStrTrim] = str[newStrTrim] + data
			} else {
				str[newStrTrim] = data
			}
			flag++
		}

	}
	answer := []string{}
	for k, v := range str {
		str := fmt.Sprintf("%v %v", v, k)
		answer = append(answer, str)
	}

	return answer
}

func main() {

	str := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	res := subdomainVisits(str)
	fmt.Println(res)
}
