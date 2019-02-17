package golang

func TwoSum(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 != k2 && v1+v2 == target {
				return []int{k1, k2}
			}
		}
	}
	return nil
}

// Once hash table

func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		temp := target - value
		if i, ok := m[temp]; ok && m[temp] != index {
			return []int{i, index}
		}
		// 键值调转排列
		m[value] = index
	}
	return nil
}
