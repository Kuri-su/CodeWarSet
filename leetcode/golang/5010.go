package golang

func digitsCount(d int, low int, high int) int {
	// 无法通过的代码...超时
	var res int
	for i := low; i <= high; i++ {
		// log.Println(i)
		tmp := i
		for {
			if d == tmp%10 {
				res++
			}
			tmp /= 10
			if tmp == 0 {
				break
			}
		}
	}
	return res
}
