package main

func moveZeroes(nums []int) {
	newArr := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			newArr = append(newArr, nums[i])
		}
	}
	for i := 0; i <= len(nums)-len(newArr); i++ {
		newArr = append(newArr, 0)
	}
	nums = newArr
}

func main() {
	nums := []int{0, 1, 0, 3, 5, 6, 7}
	moveZeroes(nums)

}
