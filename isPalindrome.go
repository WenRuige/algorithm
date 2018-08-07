package main

func isPalindrome(x int) bool {

	origin := x

	if x < 0 {
		return false
	}
	res := 0
	for x > 0 {
		pop := x % 10
		x = x / 10

		res = res*10 + pop

	}
	if res == origin {
		return true
	}
	return false
}
