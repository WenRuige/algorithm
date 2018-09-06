package main

//import "fmt"
//
//type Node struct {
//	char   rune // utf-8
//	childs map[rune]*Node
//	Data   interface{} // 数据
//	deep   int         // 树的深度
//	isTerm bool        // 一颗树是否到底
//}
//
//type Trie struct {
//	root *Node
//	size int
//}
//
//func NewNode(char rune, deep int) *Node {
//	node := &Node{
//		char:   char,
//		childs: make(map[rune]*Node, 16),
//		deep:   deep,
//	}
//	return node
//}
//
//func NewTrie() *Trie {
//	trie := &Trie{
//		root: NewNode(' ', 1),
//		size: 0,
//	}
//	return trie
//}
//
//func (t *Trie) Add(key string, data interface{}) {
//	// 遍历字符串
//	var parent *Node = t.root
//	allChars := []rune(key)
//	for _, char := range allChars {
//		node, ok := parent.childs[char]
//		// 如果节点不存在的话
//		if !ok {
//			node = NewNode(char, parent.deep+1)
//			// 插入到父节点
//			parent.childs[char] = node
//		}
//		parent = node
//		//如果存在继续向下遍历
//	}
//	// data 赋值,可以传递完整的数据
//	parent.Data = data
//	parent.isTerm = true
//}
//
//// 前缀搜索
//func (t *Trie) prefix(key string) (nodes []*Node) {
//	allChars := []rune(key)
//	var node = t.root
//	// 前缀搜索
//	for _, char := range allChars {
//		child, ok := node.childs[char]
//		if !ok {
//			return // 没找到  return
//		}
//		node = child
//	}
//
//	fmt.Printf("%v+", node)
//	var queue []*Node
//	queue = append(queue, node)
//	for len(queue) > 0 {
//		var q2 []*Node
//		for _, n := range queue {
//			if n.isTerm == true {
//				nodes = append(nodes, n)
//				continue
//			}
//
//			for _, v := range n.childs {
//				q2 = append(q2, v)
//			}
//		}
//		queue = q2
//	}
//	return
//}
//
//func main() {
//	trie := NewTrie()
//	trie.Add("哈尔滨理工大学", "哈尔滨理工大学")
//	trie.Add("哈尔滨工业大学", "哈尔滨工业大学")
//	trie.Add("哈尔滨商业大学", "哈尔滨商业大学")
//	trie.Add("哈尔滨大学", "哈尔滨大学")
//	trie.Add("黑龙江大学", "黑龙江大学")
//	trie.Add("东京大学", "东京大学")
//
//	res := trie.prefix("东")
//	fmt.Println(res)
//	for _, v := range res {
//		fmt.Println(v)
//	}
//}
