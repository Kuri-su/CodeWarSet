package contest

func twoSumLessThanK(A []int, K int) int {
	c := -1

	for k1, v1 := range A {
		for k2, v2 := range A {
			if k2 == k1 {
				continue
			}
			if v1+v2 > c && v1+v2 < K {
				c = v1 + v2
			}
		}
	}
	return c
}
