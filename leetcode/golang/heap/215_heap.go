package heap

import "container/heap"

/*
   利用 堆 处理此问题, 极简

   解题思想:
		所谓栈的 POP 实际上是把最大的元素挪到最后, 逐步修复栈

   时间复杂度: O(n)
   空间复杂度: O(1)

   过程:
*/

func findKthLargest2(nums []int, k int) int {
	var res int
	n := IntHeap(nums)
	heap.Init(n)
	for i := 0; i < k; i++ {
		heap.Pop(n)
		res = n[len(n)-1]
		n = n[:len(n)-1]
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Push(x interface{}) {
}

func (h IntHeap) Pop() interface{} {
	return nil
}
