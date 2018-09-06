package main

import "fmt"

type Trie struct {
	dict       map[string]bool
	dictPrefix map[string]bool
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	d := make(map[string]bool)
	p := make(map[string]bool)
	return Trie{d, p}

}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.dict[word] = true
	for i := 0; i < len(word); i++ {
		this.dictPrefix[word[:i]] = true
	}

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.dict[word]
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.dict[prefix] || this.dictPrefix[prefix]
}

func main() {
	obj := ConstructorTrie()
	obj.Insert("apple")
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
