package golang

import "sort"

func heightChecker(heights []int) int {
	var a []int
	a = append(a, heights...)
	sort.Ints(a)
	n := 0
	for key := range heights {
		if a[key] != heights[key] {
			n++
		}
	}
	return n
}
