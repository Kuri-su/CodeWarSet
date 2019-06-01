package golang

func fixedPoint(A []int) int {
	for key, value := range A {
		if key == value {
			return key
		}
	}
	return -1
	// .... 居然过了...
}
