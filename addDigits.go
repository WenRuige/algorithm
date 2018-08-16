package main

func addDigits(num int) int {

	if num < 10 {
		return num
	}
	other := 0
	flag := false
	for !flag {
		pop := num % 10
		num = num / 10
		other = num + pop
		num = other
		if other < 10 {
			flag = true
		}
	}

	return other

}

func main() {
	addDigits(38)
}
