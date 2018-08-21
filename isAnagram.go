package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {

	mapA := make(map[string]int)
	for i := 0; i < len(s); i++ {
		_, ok := mapA[string(s[i])]
		if ok {
			mapA[string(s[i])] = mapA[string(s[i])] + 1
		} else {
			mapA[string(s[i])] = 1
		}
	}

	mapB := make(map[string]int)
	for i := 0; i < len(t); i++ {
		_, ok := mapA[string(t[i])]
		if ok {
			mapB[string(t[i])] = mapB[string(t[i])] + 1
		} else {
			mapB[string(t[i])] = 1
		}
	}

	if reflect.DeepEqual(mapA, mapB) {
		return true
	}

	//data1 := 0
	//data2 := 0
	//if len(s) != len(t) {
	//	return false
	//}
	//
	//for i := 0; i < len(s); i++ {
	//	data1 = data1 + int(s[i])
	//}
	//
	//for i := 0; i < len(t); i++ {
	//	data2 = data2 + int(t[i])
	//}
	//
	//return data1 == data2

	return false
}

func main() {
	res := isAnagram("ABC", "ABC")
	fmt.Println(res)
}
