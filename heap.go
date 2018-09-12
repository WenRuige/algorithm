package main

import "container/heap"
import "fmt"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	//h := &IntHeap{2, 1, 5}
	//heap.Init(h)
	//heap.Push(h, 3)
	//fmt.Printf("minimum: %d\n", (*h)[0])
	//for h.Len() > 0 {
	//	fmt.Printf("%d \n", heap.Pop(h))
	//}
	nums := []int{3, 2, 1, 5, 6, 4}
	res := findKthLargest(nums, 2)
	fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)

	for h.Len() > k {
		heap.Pop(&h)
	}

	return heap.Pop(&h).(int)
}
