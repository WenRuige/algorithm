package main

import (
	"sort"
)

type strStruct struct {
	word      string
	frequence int
}
type freWords []*strStruct

func (f freWords) Len() int {
	return len(f)

}

// 进行交换
func (f freWords) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f freWords) Less(i, j int) bool {
	// 如果连个词出现的次数相同的时候,需要判断word的长度
	if f[i].frequence == f[j].frequence {
		return f[i].word < f[j].word
	}
	return f[i].frequence > f[j].frequence
}

func topKFrequent2(words []string, k int) []string {
	count := make(map[string]int, len(words))

	for _, w := range words {
		count[w]++
	}

	// 将数据都解析到struct上
	fw := make(freWords, 0, len(words))
	for w, c := range count {
		fw = append(fw, &strStruct{
			word:      w,
			frequence: c,
		})
	}
	sort.Sort(fw)

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = fw[i].word
	}

	return res
}

func main() {
	topKFrequent2([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2)
}
