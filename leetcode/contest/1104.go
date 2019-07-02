package contest

func distributeCandies(candies int, num_people int) []int {
	r := make([]int, num_people)

	n := 1
	for i := 0; i < len(r); i++ {
		if n < candies {
			r[i] += n
			candies -= n
			n++
		} else {
			r[i] += candies
			break
		}
		if i == len(r)-1 {
			if n != 0 {
				i = -1
			}
		}
	}

	return r
}
