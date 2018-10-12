package main

import "fmt"

func subsets(nums []int) [][]int {
	empty := []int{}
	sets := [][]int{empty}
	for _, n := range nums {
		newSets := make([][]int, 0)
		for _, set := range sets {
			tmp := make([]int, len(set)+1)
			copy(tmp, append(set, n))
			newSets = append(newSets, tmp)
		}
		sets = append(sets, newSets...)
		fmt.Println(sets)
	}
	return sets
}

func main() {
	subsets([]int{1, 2, 3})
}
