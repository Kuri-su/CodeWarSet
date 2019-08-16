package heap

import "container/heap"

/*
   利用 堆 处理此问题, 极简

   解题思想:

   时间复杂度: O(logN)
   空间复杂度: O(n) (可优化)

   过程:
*/

func findKthLargest2(nums []int, k int) int {

}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return i < j
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Push(x interface{}) {

}

func (h IntHeap) Pop() interface{} {
	return nil
}
