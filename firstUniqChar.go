package main

import "fmt"

//func firstUniqChar(s string) int {
//
//	for i := 0; i < len(s); i++ {
//		flag := false
//		for j := 0; j < len(s); j++ {
//			if string(s[i]) == string(s[j]) && i != j {
//				flag = true
//			}
//		}
//
//		if !flag {
//			return i
//		} else {
//			flag = false
//		}
//	}
//	return -1
//}

func firstUniqChar(s string) int {
	arr := make([]int, 26)
	for i := 0; i < 26; i++ {
		arr[i] = 0
	}
	for i := 0; i < len(s); i++ {
		bytes := int(s[i] - 'a')
		arr[bytes]++
	}

	fmt.Printf("%v", arr)

	for i := 0; i < len(s); i++ {
		bytes := int(s[i] - 'a')
		if arr[bytes] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	res := firstUniqChar("loveleetcode")
	fmt.Println(res)
}
