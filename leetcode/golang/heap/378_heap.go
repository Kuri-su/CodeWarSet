package heap

import (
	"container/heap"
)

type MyHeap1 []int

func (m MyHeap1) Len() int {
	return len(m)
}

func (m MyHeap1) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyHeap1) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyHeap1) Push(x interface{}) {

}

func (m MyHeap1) Pop() interface{} {
	return nil
}

func kthSmallest(matrix [][]int, k int) int {
	mm := MyHeap1{}
	// 构建堆
	for _, value := range matrix {
		mm = append(mm, value...)
	}
	// 取出第k小的数
	var res int
	heap.Init(mm)
	for i := 0; i < k; i++ {
		heap.Pop(mm)
		res = mm[len(mm)-1]
		mm = mm[:len(mm)-1]
	}

	return res
}
