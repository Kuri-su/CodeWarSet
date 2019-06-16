package golang

// 暴力解法

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
