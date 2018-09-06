package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
		"0": []string{" "},
	}
	slices := [][]string{}
	for index := range digits {
		digit := string(digits[index])
		if _, ok := m[digit]; ok {
			slices = append(slices, m[digit])
		}
	}
	fmt.Println(slices)
	return generateCombinations(slices)
}

func generateCombinations(slices [][]string) (ret []string) {
	if len(slices) == 0 {
		return
	}
	if len(slices) == 1 {
		return slices[0]
	}
	for _, letter := range slices[0] {
		//fmt.Println(letter)
		tmpRets := generateCombinations(slices[1:])
		//fmt.Println(slices[1:])

		for _, tmp := range tmpRets {
			ret = append(ret, letter+tmp)
		}
	}
	return
}

//func letterCombinations(digits string) []string {
//	data := make(map[string]string)
//	data["2"] = "abc"
//	data["3"] = "def"
//	data["4"] = "ghi"
//	data["5"] = "jkl"
//	data["6"] = "mno"
//	data["7"] = "pqrs"
//	data["8"] = "tuv"
//	data["9"] = "wxyz"
//	str := ""
//	tempArr := []string{}
//	newStr := []string{}
//	for i := 0; i < len(digits); i++ {
//		str = str + data[string(digits[i])]
//		newStr = append(newStr, data[string(digits[i])])
//	}
//
//	if len(digits) == 1 {
//		for i := 0; i < len(str); i++ {
//			tempArr = append(tempArr, string(str[i]))
//		}
//		return tempArr
//	}
//
//	res := make(map[string]int)
//	for i := 0; i < len(str); i++ {
//		for j := i; j < len(str); j++ {
//			//fmt.Println(string(str[i]) + string(str[j]))
//			if string(str[i]) == string(str[j]) {
//				continue
//			} else {
//				_, ok := res[string(str[i])+string(str[j])]
//				if ok {
//					res[string(str[i])+string(str[j])] = 1
//				} else {
//					res[string(str[i])+string(str[j])] = res[string(str[i])+string(str[j])] + 1
//				}
//			}
//		}
//	}
//
//	for k := 0; k < len(newStr); k++ {
//		for i := 0; i < len(newStr[k]); i++ {
//			for j := i; j < len(newStr[k]); j++ {
//				//fmt.Println(string(str[i]) + string(str[j]))
//				if string(newStr[k][i]) == string(newStr[k][j]) {
//					continue
//				} else {
//
//					_, ok := res[string(newStr[k][i])+string(newStr[k][j])]
//					if ok {
//						delete(res, string(newStr[k][i])+string(newStr[k][j]))
//					}
//
//				}
//			}
//		}
//	}
//
//	for k, _ := range res {
//		tempArr = append(tempArr, k)
//	}
//	return tempArr
//}

func main() {
	res := letterCombinations("234")
	fmt.Println(res)
}
