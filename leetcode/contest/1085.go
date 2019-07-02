package contest

import "strconv"

func sumOfDigits(A []int) int {
	min := -1
	for _, value := range A {
		if min == -1 {
			min = value
		}
		if value < min {
			min = value
		}
	}

	itoa := strconv.Itoa(min)
	sum := 0
	for _, value := range itoa {
		i, _ := strconv.Atoi(string(value))
		sum += i
	}
	if sum%2 == 1 {
		return 0
	} else {
		return 1
	}
}
