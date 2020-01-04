package toSimple

/*
  解题思想:
	暴力

   时间复杂度: O(n)
   空间复杂度: O(1)
*/

func findMaxConsecutiveOnes(nums []int) int {
	var max, onceMax int

	for _, v := range nums {
		switch v {
		case 1:
			onceMax++
		case 0:
			if max < onceMax {
				max = onceMax
			}
			onceMax = 0
		}
	}
	if max < onceMax {
		max = onceMax
	}
	onceMax = 0
	return max
}
