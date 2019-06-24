package golang

// 暴力解法
// 按层找水
/*
   这个应该很好理解,

     = 0 0 =    = 代表柱子
     = 0 0 =    0 代表雨水
     - - - -    - 代表地面
     0 1 2 3

   解法思想：碰到比当前高的柱子,就看看右,是否有大于或等于当前高度的柱子, 使得可以接住水

   时间复杂度: O(n^2)
   空间复杂度: O(1)

   这个解法比较直观,所以详细的步骤省略啦~
*/

func trap(height []int) int {
	var max = -1
	for _, value := range height {
		if value > max {
			max = value
		}
	}

	result := 0

	for i := 1; i <= max; i++ {
		for j := 0; j < len(height); j++ {
			if height[j] >= i {
				// 往右看, 直到找到和当前高度相等或者更高的高度
				for k := j + 1; k < len(height); k++ {
					if height[k] >= i {
						result += k - (j + 1)
						j = k
					}
				}
			}
		}
	}

	return result
}
