package main

import (
	"fmt"
)

//func addBinary(a string, b string) string {
//
//	if a == "0" && b == "0" {
//		return "0"
//	}
//	aSum := 0
//	bSum := 0
//	k := 0
//	for i := len(a) - 1; i >= 0; i-- {
//		intNum, _ := strconv.Atoi(string(a[i]))
//		aSum = aSum + int(math.Pow(2, float64(k)))*intNum
//		k++
//	}
//	k = 0
//	for i := len(b) - 1; i >= 0; i-- {
//		intNum, _ := strconv.Atoi(string(b[i]))
//		bSum = bSum + int(math.Pow(2, float64(k)))*intNum
//		k++
//	}
//
//	total := aSum + bSum
//
//	fmt.Println(total)
//	data := []int{}
//	for total > 0 {
//		pop := total % 2
//		total = total / 2
//		data = append(data, pop)
//	}
//
//	data2 := ""
//
//	for i := len(data) - 1; i >= 0; i-- {
//		data2 = data2 + strconv.Itoa(data[i])
//	}
//
//	return data2
//}

func addBinary(a string, b string) string {

	return ""
}

func main() {
	res := addBinary("1010", "1011")
	fmt.Println(res)
}
