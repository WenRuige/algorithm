package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {

		if g[i] <= s[j] {
			i++
			j++
			count++
		} else {
			j++
		}
	}
	return count

}

func main() {
	g := []int{1, 1}
	s := []int{1, 4, 3}
	res := findContentChildren(g, s)
	fmt.Println(res)
}
