package contest

func rearrangeBarcodes(barcodes []int) []int {
	a := map[int]int{}
	for _, value := range barcodes {
		if _, ok := a[value]; !ok {
			a[value] = 1
		} else {
			a[value]++
		}
	}

	b := make([]int, 0, 10000)

	for {
		flag := true
		for key, value := range a {
			if value > 0 {
				b = append(b, key)
				a[key]--
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
	return b
}
