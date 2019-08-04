package contest

func movesToMakeZigzag(nums []int) int {
	// 满足第一种的 操作次数
	nums2 := make([]int, 0, len(nums))
	nums2 = append(nums2, nums...)

	total1 := 0
	for i := 1; i < len(nums); i += 2 {
		// 前
		if nums[i] >= nums[i-1] {
			total1 += nums[i] - nums[i-1] + 1
			nums[i] = nums[i-1] - 1
		}

		// 后
		if i+1 < len(nums) && nums[i] >= nums[i+1] {
			total1 += nums[i] - nums[i+1] + 1
			nums[i] = nums[i+1] - 1
		}
	}

	// 满足第二种的操作次数
	total2 := 0
	for i := 0; i < len(nums2); i += 2 {
		// 前
		if i-1 > 0 && nums2[i] >= nums2[i-1] {
			total2 += nums2[i] - nums2[i-1] + 1
			nums2[i] = nums2[i-1] - 1
		}

		// 后
		if i+1 < len(nums2) && nums2[i] >= nums2[i+1] {
			total2 += nums2[i] - nums2[i+1] + 1
			nums2[i] = nums2[i+1] - 1
		}
	}

	// log.Println(total1, "  ", total2)

	if total1 <= total2 {
		return total1
	} else {
		return total2
	}

}
