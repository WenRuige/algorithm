package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	strArrA := strings.Split(A, " ")
	strArrB := strings.Split(B, " ")
	strMap := make(map[string]int)
	for i := 0; i < len(strArrA); i++ {
		_, ok := strMap[strArrA[i]]
		if ok {

			strMap[strArrA[i]] = strMap[strArrA[i]] + 1
		} else {
			strMap[strArrA[i]] = 1
		}
	}
	for i := 0; i < len(strArrB); i++ {
		_, ok := strMap[strArrB[i]]
		if ok {
			strMap[strArrB[i]] = strMap[strArrB[i]] + 1
		} else {
			strMap[strArrB[i]] = 1
		}
	}
	str := []string{}
	for k, v := range strMap {
		//fmt.Println(k, v)
		if v == 1 {
			str = append(str, k)
		}
	}

	return str
}

func main() {
	res := uncommonFromSentences("this apple is sweet", "this apple is sour")
	fmt.Println(res)
}
