package main

import "fmt"

// 692
func topKFrequent2(words []string, k int) []string {

	str := make(map[string]int)
	for i := 0; i < len(words); i++ {
		_, ok := str[words[i]]
		if ok {
			str[words[i]] = str[words[i]] + 1
		} else {
			str[words[i]] = 1
		}
	}

	// 找出前两个元素

	fmt.Println(str)



	return []string{}
}

func main(){
	topKFrequent2([]string{"i", "love", "leetcode", "i", "love", "coding"},2)
}