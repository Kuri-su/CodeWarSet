package toSimple

import (
	"math"
)

func findNumbers(nums []int) int {
	var sum int
	for _, v := range nums {
		times := int(math.Floor(math.Log10(float64(v)))) + 1
		if times != 0 && times%2 == 0 {
			sum++
		}
	}
	return sum
}
