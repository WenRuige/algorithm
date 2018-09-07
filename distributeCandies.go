package main

import (
	"fmt"
)

func distributeCandies(candies []int) int {
	kind := make(map[int]int)
	for i := 0; i < len(candies); i++ {
		_, ok := kind[candies[i]]
		if ok {
			kind[candies[i]] = kind[candies[i]] + 1
		} else {
			kind[candies[i]] = 1
		}
	}

	fmt.Println(kind)

	if len(kind) > (len(candies) / 2) {
		return (len(candies) / 2)
	} else {
		return len(kind)
	}
}

func main() {
	candies := []int{1, 2, 3, 4, 5, 6}
	res := distributeCandies(candies)
	fmt.Println(res)
}
