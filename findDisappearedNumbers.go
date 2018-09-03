package main

func findDisappearedNumbers(nums []int) []int {
	//sort.Ints(nums)
	newArr := []int{}
	for i := 0; i < len(nums); i++ {
		id := abs(nums[i]) - 1
		if nums[id] > 0 {
			nums[id] = -nums[id]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			newArr = append(newArr, i+1)
		}
	}

	return newArr

}
func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDisappearedNumbers(nums)
}
