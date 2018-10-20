package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	path := make([]int, len(graph))
	res := make([][]int, 0, len(graph))
	// 长度减一代表最后一个元素一定是空节点
	dfs(0, len(graph)-1, 1, path, graph, &res)

	return res
}

func dfs(id, dst, pathLen int, path []int, graph [][]int, res *[][]int) {
	if id == dst {
		temp := make([]int, pathLen)
		copy(temp, path)
		*res = append(*res, temp)
	}

	for _, node := range graph[id] {
		path[pathLen] = node
		dfs(node, dst, pathLen+1, path, graph, res)
	}
}

func main() {
	nums := [][]int{{1, 2}, {3}, {3}, {}}
	res := allPathsSourceTarget(nums)
	fmt.Println(res)
}
