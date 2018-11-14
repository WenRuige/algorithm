package main

import (
	"fmt"
	"sort"
	"strings"
)

// trie 前缀树
func replaceWords(dict []string, sentence string) string {
	hasRoot := make(map[string]bool, len(dict))
	hasLen := make([]bool, 101)
	for i := range dict {
		hasRoot[dict[i]] = true
		hasLen[len(dict[i])] = true
	}

	ls := make([]int, 0, 101)
	for i, ok := range hasLen {
		if ok {
			ls = append(ls, i)
		}
	}
	sort.Ints(ls)
	words := strings.Split(sentence, " ")
	for i, w := range words {
		for _, l := range ls {
			//fmt.Println(l)
			// w 的前 l 个字符是字根就修改 w 并结束 for 循环
			if l < len(w) && hasRoot[w[:l]] {
				words[i] = w[:l]
				break
			}
		}
	}

	return strings.Join(words, " ")
}

func main() {
	res := replaceWords([]string{"cats", "bat", "rat"}, "the cattle was rattled by the battery")
	fmt.Println(res)
}
