package main

import (
	"fmt"
	"strconv"
)

// 在看看
func isIsomorphic(s string, t string) bool {
	// 长度不一样直接返回
	if len(s) != len(t) {
		return false
	}
	//

	structA := make(map[string]int)
	structB := make(map[string]int)
	flagA := 0
	flagB := 0
	strA := ""
	strB := ""
	for i := 0; i < len(s); i++ {
		_, ok := structA[string(s[i])]
		if ok {
			strA = strA + strconv.Itoa(structA[string(s[i])])
		} else {
			structA[string(s[i])] = flagA
			strA = strA + strconv.Itoa(flagA)
			flagA++

		}

	}

	for i := 0; i < len(t); i++ {
		_, ok := structB[string(t[i])]
		if ok {
			strB = strB + strconv.Itoa(structB[string(t[i])])
		} else {
			structB[string(t[i])] = flagB
			strB = strB + strconv.Itoa(flagB)
			flagB++

		}

	}

	if strA == strB {
		return true
	}

	fmt.Println(strA, strB)
	return false
}

func main() {
	res := isIsomorphic("egg", "add")
	fmt.Println(res)
}
