package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	p1 := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		_, ok := p1[nums1[i]]
		if ok {
			p1[nums1[i]] = p1[nums1[i]] + 1
		} else {
			p1[nums1[i]] = 1
		}
	}
	total := []int{}
	for i := 0; i < len(nums2); i++ {
		_, ok := p1[nums2[i]]
		if ok {
			//total = append(total, nums2[i])
			if p1[nums2[i]] >= 1 {
				total = append(total, nums2[i])
				p1[nums2[i]] = p1[nums2[i]] - 1
			} else {
				delete(p1, nums2[i])
			}

		}
	}

	return total

}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}
