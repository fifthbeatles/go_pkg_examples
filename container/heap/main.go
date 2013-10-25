package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Print() {
	for _, v := range h {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	arr := []int{2, 1, 5, 9, 4, 6, 8, 0, 7, 3}
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range arr {
		heap.Push(h, v)
		h.Print()
	}
	heap.Remove(h, 1)
	h.Print()
	_ = heap.Pop(h)
	h.Print()
}
