/*
Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:

Rank is an integer starting from 1.
The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
Rank should be as small as possible.
*/

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}

func arrayRankTransform(arr []int) []int {
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range arr {
		heap.Push(h, v)
	}

	m := make(map[int]int)

	cntr := 1
	for h.Len() > 0 {
		v := heap.Pop(h)
		if _, ok := m[v.(int)]; !ok {
			m[v.(int)] = cntr
			cntr++
		}
	}

	ret := make([]int, len(arr))
	for i, v := range arr {
		ret[i] = m[v]
	}
	return ret

}
