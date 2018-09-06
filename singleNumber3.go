package main

func singleNumber3(nums []int) int {

	if len(nums) <= 1 {
		return nums[0]
	}
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = hashmap[nums[i]] + 1
	}

	for v := range hashmap {
		if hashmap[v] == 1 {
			return v

		}
	}
	return 0
}

func main() {
	//singleNumber()
}
