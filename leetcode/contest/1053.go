package contest

import "sort"

func prevPermOpt1(A []int) []int {
	var a []int
	a = append(a, A...)
	sort.Ints(a)

	flag := false
	for key := range A {
		if a[key] != A[key] {
			flag = true
			break
		}
	}

	if flag {

		max := a[len(a)-1]

		n := -1
		for key, value := range A {
			if value == max {
				n = key
			}
		}
		A[n], A[len(a)] = A[len(a)], max

	}
	return A
}
