package main

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// 添加Node
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(u, v *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 第一次创建,没边,分配内存
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	// 因为是无向图
	g.edges[*u] = append(g.edges[*u], v)
	g.edges[*v] = append(g.edges[*v], u)

}

func (g *Graph) String() {
	g.lock.Lock()
	defer g.lock.Unlock()

	str := ""
	for _, iNode := range g.nodes {
		str += iNode.String() + " -> "
		nexts := g.edges[*iNode]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func main() {

	g := Graph{}
	n1, n2, n3, n4, n5 := Node{1}, Node{2}, Node{3}, Node{4}, Node{5}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	//g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)

	g.String()
}
