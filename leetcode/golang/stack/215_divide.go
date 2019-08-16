package stack

/*
   利用 分治的思想 处理此问题, 具体表现为 快排 + 类似二分查找

   解题思想:
		主体思想是分治,但是如何分治呢, 如果直接分成十块来分治, 那么将很容易陷入局部最优.
		如果使用快排的思想, 取一根中轴,然后将数组中的值分到两边的做法, 然后确定在哪一半, 然后在有的那一半再继续取中轴, 如此重复下去即可, 看起来有点像 快排 + 二分. 这样可以在避免陷入局部最优的前提下,
以尽少的时间复杂度找到结果.

   时间复杂度: O(logN)
   空间复杂度: O(n) (可优化)

   过程:
*/

func findKthLargest(nums []int, k int) int {
	for {
		// 取第 K 大的元素
		left, right, axis := block(nums)
		if len(right)+1 == k {
			return axis
		} else if len(right) >= k {
			nums = right
		} else {
			if len(left) == 1 && len(right)+1+1 == k {
				return left[0]
			}
			nums = left
			k = k - len(right) - 1
		}
	}
}

func block(arr []int) ([]int, []int, int) {
	axis := arr[len(arr)-1] // 轴
	leftArr := []int{}      // 小于或者等于axis
	rightArr := []int{}     // 大于axis

	for key, value := range arr {
		if key == len(arr)-1 {
			continue
		}
		if value <= axis {
			leftArr = append(leftArr, value)
		} else {
			rightArr = append(rightArr, value)
		}
	}

	return leftArr, rightArr, axis
}
