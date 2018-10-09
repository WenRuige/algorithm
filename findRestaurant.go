package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {

	newArr := map[string]int{}
	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				newArr[list1[i]] = i + j
			}
		}
	}

	// 找出最小值
	newStrArr := []string{}
	min := 1 << 32
	for k, v := range newArr {
		if min > v {
			newStrArr = []string{}
			min = v
			newStrArr = append(newStrArr, k)
		} else if min == v {
			newStrArr = append(newStrArr, k)
		}
	}
	return newStrArr
}

func main() {
	str1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	str2 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(str1, str2))
}
